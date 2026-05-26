package orm

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) CreateCart(id int64) (*domain.Cart, error) {
	var cart = &domain.Cart{
		UserID: &id,
		Status: "active",
	}
	if err := r.db.Create(cart).Error; err != nil {
		return nil, ports.ErrFailedToSaveCart

	}
	return cart, nil
}

func (r *CartRepository) GetCartWithItems(userID int64) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.db.Preload("Items.Product").Where("user_id = ? AND status = ?", userID, "active").First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.CreateCart(userID)
		}
		return nil, ports.ErrFailedToLoadCart
	}
	for i := range cart.Items {
		if cart.Items[i].PriceSnapshot <= 0 && cart.Items[i].Product != nil {
			cart.Items[i].PriceSnapshot = cart.Items[i].Product.Price
		}
	}
	return &cart, nil
}

func (r *CartRepository) FindCartItem(cartID, productID int64) (*domain.CartItem, error) {
	var item domain.CartItem

	err := r.db.
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		First(&item).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrCartItemNotFound
		}
		return nil, ports.ErrFailedToLoadCartItem
	}

	return &item, nil
}

func (r *CartRepository) DeleteCartItem(userId int64, itemId int64) error {
	var cart domain.Cart

	err := r.db.
		Where("user_id = ? AND status = ?", userId, "active").
		First(&cart).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ports.ErrCartNotFound
		}
		return ports.ErrFailedToLoadCart
	}

	res := r.db.Where("cart_id = ? AND id = ?", cart.ID, itemId).Delete(&domain.CartItem{})
	if res.Error != nil {
		return ports.ErrFailedToDeleteCartItem
	}

	if res.RowsAffected == 0 {
		return ports.ErrCartItemNotFound
	}
	return nil
}

func (r *CartRepository) UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error) {
	var item domain.CartItem
	if err := r.db.First(&item, itemId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrCartItemNotFound
		}
		return nil, ports.ErrFailedToLoadCart
	}
	item.Quantity = quantity
	if err := r.db.Save(&item).Error; err != nil {
		return nil, ports.ErrFailedToUpdateCartItem
	}
	return &item, nil
}

func (r *CartRepository) AddCartItem(userId int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	var cart domain.Cart
	err := r.db.
		Where("user_id = ? AND status = ?", userId, "active").
		First(&cart).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart = domain.Cart{
				UserID: &userId,
				Status: "active",
			}

			if err := r.db.Create(&cart).Error; err != nil {
				return nil, ports.ErrFailedToSaveCart
			}
		} else {
			return nil, ports.ErrFailedToLoadCart
		}
	}

	var existingItem domain.CartItem
	err = r.db.
		Where("cart_id = ? AND product_id = ?", cart.ID, cartItem.ProductID).
		First(&existingItem).Error

	if err == nil {
		existingItem.Quantity += cartItem.Quantity
		if existingItem.PriceSnapshot <= 0 && cartItem.PriceSnapshot > 0 {
			existingItem.PriceSnapshot = cartItem.PriceSnapshot
		}

		if err := r.db.Save(&existingItem).Error; err != nil {
			return nil, ports.ErrFailedToUpdateCartItem
		}

		return &existingItem, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ports.ErrFailedToLoadCart
	}

	cartItem.CartID = cart.ID

	if err := r.db.Create(cartItem).Error; err != nil {
		return nil, ports.ErrFailedToSaveCart
	}

	return cartItem, nil
}

func (r *CartRepository) ClearCart(userId int64) error {
	var cart domain.Cart
	err := r.db.Where("user_id = ? AND status = ?", userId, "active").First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ports.ErrCartNotFound
		}
		return ports.ErrFailedToClearCart
	}
	if err = r.db.Delete(&domain.Cart{}, cart.ID).Error; err != nil {
		return ports.ErrFailedToClearCart
	}
	return nil
}
