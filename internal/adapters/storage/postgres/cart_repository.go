package postgres

import (
	"context"
	"database/sql"
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
	"time"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) CreateCart(id int64) (*domain.Cart, error) {
	var cart = &domain.Cart{
		UserID: &id,
		Status: "active",
	}
	query := `INSERT INTO carts(user_id, status, created_at, updated_at) 
          VALUES ($1, 'active', NOW(), NOW()) 
          RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, id).Scan(&cart.ID, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, ports.ErrFailedToSaveCart
	}
	return cart, nil
}

func (r *CartRepository) GetCartWithItems(userID int64) (*domain.Cart, error) {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, ports.ErrFailedToLoadCart
	}
	defer tx.Rollback()

	var cartID int64
	var cartStatus string
	var createdAt, updatedAt time.Time

	err = tx.QueryRowContext(ctx, `
        SELECT id, status, created_at, updated_at
        FROM carts 
        WHERE user_id = $1 AND status = 'active'
        FOR UPDATE
    `, userID).Scan(&cartID, &cartStatus, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		err = tx.QueryRowContext(ctx, `
            INSERT INTO carts (user_id, status, created_at, updated_at)
            VALUES ($1, 'active', NOW(), NOW())
            RETURNING id, status, created_at, updated_at
        `, userID).Scan(&cartID, &cartStatus, &createdAt, &updatedAt)

		if err != nil {
			return nil, ports.ErrFailedToLoadCart
		}
	} else if err != nil {
		return nil, ports.ErrFailedToLoadCart
	}

	rows, err := tx.QueryContext(ctx, `
        SELECT 
            ci.id,
            ci.product_id,
            ci.quantity,
            ci.price_snapshot,
            p.id,
            p.name,
            p.price,
            COALESCE(p.description, '')
        FROM cart_items ci
        LEFT JOIN products p ON p.id = ci.product_id
        WHERE ci.cart_id = $1
        ORDER BY ci.id
    `, cartID)

	if err != nil {
		return nil, ports.ErrFailedToLoadCart
	}
	defer rows.Close()

	cart := &domain.Cart{
		ID:        cartID,
		UserID:    &userID,
		Status:    cartStatus,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Items:     []domain.CartItem{},
	}

	for rows.Next() {
		var item domain.CartItem
		var product domain.Product
		var priceSnapshot sql.NullFloat64

		err := rows.Scan(
			&item.ID, &item.ProductID, &item.Quantity, &priceSnapshot,
			&product.ID, &product.Name, &product.Price, &product.Description,
		)
		if err != nil {
			return nil, ports.ErrFailedToLoadCart
		}

		if priceSnapshot.Valid && priceSnapshot.Float64 > 0 {
			item.PriceSnapshot = priceSnapshot.Float64
		} else {
			item.PriceSnapshot = product.Price
		}

		item.Product = &product
		cart.Items = append(cart.Items, item)
	}
	if err := tx.Commit(); err != nil {
		return nil, ports.ErrFailedToLoadCart
	}

	return cart, nil
}

func (r *CartRepository) FindCartItem(cartID, productID int64) (*domain.CartItem, error) {
	var item domain.CartItem

	query := `SELECT id, cart_id, product_id, quantity, price_snapshot, created_at, updated_at 
              FROM cart_items 
              WHERE cart_id = $1 AND product_id = $2`

	err := r.db.QueryRow(query, cartID, productID).Scan(
		&item.ID, &item.CartID, &item.ProductID,
		&item.Quantity, &item.PriceSnapshot,
		&item.CreatedAt, &item.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, ports.ErrCartItemNotFound
	}
	if err != nil {
		return nil, ports.ErrFailedToLoadCartItem
	}
	return &item, nil
}

func (r *CartRepository) DeleteCartItem(userId int64, itemId int64) error {
	var cartId int64
	err := r.db.QueryRow(`
        SELECT id FROM carts 
        WHERE user_id = $1 AND status = 'active'
    `, userId).Scan(&cartId)

	if err == sql.ErrNoRows {
		return ports.ErrCartNotFound
	}
	if err != nil {
		return ports.ErrFailedToLoadCart
	}

	result, err := r.db.Exec(`
        DELETE FROM cart_items 
        WHERE cart_id = $1 AND id = $2
    `, cartId, itemId)

	if err != nil {
		return ports.ErrFailedToDeleteCartItem
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return ports.ErrFailedToDeleteCartItem
	}

	if rows == 0 {
		return ports.ErrCartItemNotFound
	}

	return nil
}

func (r *CartRepository) UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error) {
	var item domain.CartItem
	query := `UPDATE cart_items 
          SET quantity = $1, updated_at = NOW() 
          WHERE id = $2 
          RETURNING id, cart_id, product_id, quantity, price_snapshot, created_at, updated_at`
	row := r.db.QueryRow(query, quantity, itemId)
	err := row.Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity,
		&item.PriceSnapshot, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, ports.ErrCartItemNotFound
		}
		return nil, ports.ErrFailedToUpdateCartItem
	}
	return &item, nil
}

func (r *CartRepository) AddCartItem(userID int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, ports.ErrFailedToSaveCart
	}
	defer tx.Rollback()

	var cartID int64
	err = tx.QueryRowContext(ctx, `
        SELECT id FROM carts 
        WHERE user_id = $1 AND status = 'active'
        FOR UPDATE
    `, userID).Scan(&cartID)

	if err == sql.ErrNoRows {
		err = tx.QueryRowContext(ctx, `
            INSERT INTO carts (user_id, status, created_at, updated_at)
            VALUES ($1, 'active', NOW(), NOW())
            RETURNING id
        `, userID).Scan(&cartID)

		if err != nil {
			return nil, ports.ErrFailedToSaveCart
		}
	} else if err != nil {
		return nil, ports.ErrFailedToLoadCart
	}

	var existingItem domain.CartItem
	err = tx.QueryRowContext(ctx, `
        SELECT id, quantity, price_snapshot
        FROM cart_items
        WHERE cart_id = $1 AND product_id = $2
        FOR UPDATE
    `, cartID, cartItem.ProductID).Scan(&existingItem.ID, &existingItem.Quantity, &existingItem.PriceSnapshot)

	if err == nil {
		newQuantity := existingItem.Quantity + cartItem.Quantity
		priceSnapshot := existingItem.PriceSnapshot
		if priceSnapshot <= 0 && cartItem.PriceSnapshot > 0 {
			priceSnapshot = cartItem.PriceSnapshot
		}

		err = tx.QueryRowContext(ctx, `
            UPDATE cart_items 
            SET quantity = $1, price_snapshot = $2, updated_at = NOW()
            WHERE id = $3
            RETURNING id, cart_id, product_id, quantity, price_snapshot, created_at, updated_at
        `, newQuantity, priceSnapshot, existingItem.ID).Scan(
			&cartItem.ID, &cartItem.CartID, &cartItem.ProductID,
			&cartItem.Quantity, &cartItem.PriceSnapshot,
			&cartItem.CreatedAt, &cartItem.UpdatedAt,
		)

		if err != nil {
			return nil, ports.ErrFailedToUpdateCartItem
		}
	} else if err == sql.ErrNoRows {
		cartItem.CartID = cartID
		err = tx.QueryRowContext(ctx, `
            INSERT INTO cart_items (cart_id, product_id, quantity, price_snapshot, created_at, updated_at)
            VALUES ($1, $2, $3, $4, NOW(), NOW())
            RETURNING id, created_at, updated_at
        `, cartID, cartItem.ProductID, cartItem.Quantity, cartItem.PriceSnapshot).
			Scan(&cartItem.ID, &cartItem.CreatedAt, &cartItem.UpdatedAt)

		if err != nil {
			return nil, ports.ErrFailedToSaveCart
		}
	} else {
		return nil, ports.ErrFailedToLoadCartItem
	}

	if err := tx.Commit(); err != nil {
		return nil, ports.ErrFailedToSaveCart
	}

	return cartItem, nil
}

func (r *CartRepository) ClearCart(userID int64) error {
	query := `DELETE FROM carts 
              WHERE user_id = $1 AND status = 'active'
              RETURNING id`
	var deletedID int64
	err := r.db.QueryRow(query, userID).Scan(&deletedID)
	if err == sql.ErrNoRows {
		return ports.ErrCartNotFound
	}
	if err != nil {
		return ports.ErrFailedToClearCart
	}
	return nil
}
