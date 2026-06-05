package postgres

import (
	"context"
	"database/sql"
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
	"time"

	"github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	if order == nil {
		return nil, ports.ErrFailedToSaveOrder
	}

	query := `
		INSERT INTO orders (user_id, status, total_price, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query, order.UserID, order.Status, order.TotalPrice).
		Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, ports.ErrFailedToSaveOrder
	}

	if order.Items == nil {
		order.Items = []domain.OrderItem{}
	}
	return order, nil
}

func ensureOrderItems(order *domain.Order) {
	if order.Items == nil {
		order.Items = []domain.OrderItem{}
	}
	for i, item := range order.Items {
		if item.ImageSnapshot == "" && item.Product != nil {
			order.Items[i].ImageSnapshot = item.Product.ImageURL
		}
		if item.NameSnapshot == "" && item.Product != nil && item.Product.Name != "" {
			order.Items[i].NameSnapshot = item.Product.Name
		}
	}
}

func (r *OrderRepository) GetOrderById(ctx context.Context, id int64) (*domain.Order, error) {
	var order domain.Order
	err := r.db.QueryRowContext(ctx, `
		SELECT id, user_id, status, total_price, created_at, updated_at
		FROM orders
		WHERE id = $1
	`, id).Scan(&order.ID, &order.UserID, &order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ports.ErrOrderNotFound
		}
		return nil, ports.ErrFailedToLoadOrder
	}

	items, err := r.loadOrderItemsByOrderIDs(ctx, []int64{order.ID})
	if err != nil {
		return nil, ports.ErrFailedToLoadOrder
	}
	order.Items = items[order.ID]
	ensureOrderItems(&order)
	return &order, nil
}

func (r *OrderRepository) GetOrderByUserId(ctx context.Context, userId int64) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, status, total_price, created_at, updated_at
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userId)
	if err != nil {
		return nil, ports.ErrFailedToLoadOrder
	}
	defer rows.Close()

	orders := make([]domain.Order, 0)
	orderIDs := make([]int64, 0)

	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Status, &o.TotalPrice, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, ports.ErrFailedToLoadOrder
		}
		o.Items = []domain.OrderItem{}
		orders = append(orders, o)
		orderIDs = append(orderIDs, o.ID)
	}
	if err := rows.Err(); err != nil {
		return nil, ports.ErrFailedToLoadOrder
	}
	if len(orders) == 0 {
		return orders, nil
	}

	itemsByOrderID, err := r.loadOrderItemsByOrderIDs(ctx, orderIDs)
	if err != nil {
		return nil, ports.ErrFailedToLoadOrder
	}

	for i := range orders {
		orders[i].Items = itemsByOrderID[orders[i].ID]
		ensureOrderItems(&orders[i])
	}

	return orders, nil
}

func (r *OrderRepository) AddOrderItems(ctx context.Context, orderId int64, items []domain.OrderItem) error {
	tx, err := r.db.Begin()
	if err != nil {
		return ports.ErrFailedToSaveOrder
	}
	defer tx.Rollback()

	var exists bool
	if err := tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM orders WHERE id = $1)`, orderId).Scan(&exists); err != nil {
		return ports.ErrFailedToSaveOrder
	}
	if !exists {
		return ports.ErrOrderNotFound
	}

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO order_items (order_id, product_id, quantity, price_snapshot, name_snapshot, image_snapshot, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`)
	if err != nil {
		return ports.ErrFailedToSaveOrder
	}
	defer stmt.Close()

	for i := range items {
		items[i].OrderID = orderId
		if _, err := stmt.Exec(
			items[i].OrderID,
			items[i].ProductID,
			items[i].Quantity,
			items[i].PriceSnapshot,
			items[i].NameSnapshot,
			nullIfEmpty(items[i].ImageSnapshot),
		); err != nil {
			return ports.ErrFailedToSaveOrder
		}
	}

	if err := tx.Commit(); err != nil {
		return ports.ErrFailedToSaveOrder
	}

	return nil
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, orderId int64, status string) error {
	res, err := r.db.ExecContext(ctx, `
		UPDATE orders
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`, status, orderId)
	if err != nil {
		return ports.ErrFailedToSaveOrder
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return ports.ErrFailedToSaveOrder
	}
	if affected == 0 {
		return ports.ErrOrderNotFound
	}
	return nil
}

type orderItemRow struct {
	ID            int64
	OrderID       int64
	ProductID     int64
	Quantity      int
	PriceSnapshot float64
	NameSnapshot  string
	ImageSnapshot sql.NullString
	CreatedAt     time.Time

	ProductIDDB   sql.NullInt64
	OwnerID       sql.NullInt64
	Name          sql.NullString
	Description   sql.NullString
	Price         sql.NullFloat64
	CategoryID    sql.NullInt64
	ImageURL      sql.NullString
	Stock         sql.NullInt64
	ProductCreate sql.NullTime
}

func (r *OrderRepository) loadOrderItemsByOrderIDs(ctx context.Context, orderIDs []int64) (map[int64][]domain.OrderItem, error) {
	itemsByOrderID := make(map[int64][]domain.OrderItem, len(orderIDs))
	if len(orderIDs) == 0 {
		return itemsByOrderID, nil
	}

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			oi.id,
			oi.order_id,
			oi.product_id,
			oi.quantity,
			oi.price_snapshot,
			oi.name_snapshot,
			oi.image_snapshot,
			oi.created_at,
			p.id,
			p.owner_id,
			p.name,
			p.description,
			p.price,
			p.category_id,
			p.image_url,
			p.stock,
			p.created_at
		FROM order_items oi
		LEFT JOIN products p ON p.id = oi.product_id
		WHERE oi.order_id = ANY($1)
		ORDER BY oi.order_id, oi.id
	`, pq.Array(orderIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row orderItemRow

		if err := rows.Scan(
			&row.ID,
			&row.OrderID,
			&row.ProductID,
			&row.Quantity,
			&row.PriceSnapshot,
			&row.NameSnapshot,
			&row.ImageSnapshot,
			&row.CreatedAt,
			&row.ProductIDDB,
			&row.OwnerID,
			&row.Name,
			&row.Description,
			&row.Price,
			&row.CategoryID,
			&row.ImageURL,
			&row.Stock,
			&row.ProductCreate,
		); err != nil {
			return nil, err
		}

		item := domain.OrderItem{
			ID:            row.ID,
			OrderID:       row.OrderID,
			ProductID:     row.ProductID,
			Quantity:      row.Quantity,
			PriceSnapshot: row.PriceSnapshot,
			NameSnapshot:  row.NameSnapshot,
			CreatedAt:     row.CreatedAt,
		}

		if row.ImageSnapshot.Valid {
			item.ImageSnapshot = row.ImageSnapshot.String
		}

		if row.ProductIDDB.Valid {
			item.Product = &domain.Product{
				ID:          row.ProductIDDB.Int64,
				OwnerID:     row.OwnerID.Int64,
				Name:        row.Name.String,
				Description: row.Description.String,
				Price:       row.Price.Float64,
				CategoryID:  row.CategoryID.Int64,
				ImageURL:    row.ImageURL.String,
				Stock:       int(row.Stock.Int64),
			}

			if row.ProductCreate.Valid {
				item.Product.CreatedAt = row.ProductCreate.Time
			}
		}

		itemsByOrderID[item.OrderID] = append(
			itemsByOrderID[item.OrderID],
			item,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return itemsByOrderID, nil
}

func nullIfEmpty(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}
