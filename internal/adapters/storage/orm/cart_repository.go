package orm

import (
	"errors"
	"market/internal/core/domain"

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
	res := r.db.Create(cart)
	if res.Error != nil {
		return nil, errors.New(res.Error.Error())
	}
	return cart, nil
}

func (r *CartRepository) GetCartWithItems(userID int64) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.db.Preload("Items.Product").Where("user_id = ? AND status = ?", userID, "active").First(&cart).Error
	if err != nil {
		cart, err := r.CreateCart(userID)
		if err != nil {
			return nil, err
		}
		return cart, nil
	}
	return &cart, nil
}

func (r *CartRepository) FindCartItem(cartID, productID int64) (*domain.CartItem, error) {
	var item domain.CartItem

	err := r.db.
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		First(&item).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *CartRepository) DeleteCartItem(userId int64, itemId int64) error {
	var cart domain.Cart

	err := r.db.
		Where("user_id = ? AND status = ?", userId, "active").
		First(&cart).Error

	if err != nil {
		return err
	}

	return r.db.
		Where("cart_id = ? AND id = ?", cart.ID, itemId).
		Delete(&domain.CartItem{}).Error
}

func (r *CartRepository) UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error) {
	var item domain.CartItem
	res := r.db.Find(&item, itemId)
	if res.Error != nil {
		return nil, res.Error
	}
	item.Quantity = quantity
	err := r.db.Save(&item).Error
	if err != nil {
		return &item, err
	}
	return &item, nil
}

func (r *CartRepository) AddCartItem(userId int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	var cart domain.Cart
	err := r.db.
		Where("user_id = ? AND status = ?", userId, "active").
		First(&cart).Error

	if err != nil {
		cart = domain.Cart{
			UserID: &userId,
			Status: "active",
		}

		if err := r.db.Create(&cart).Error; err != nil {
			return nil, err
		}
	}

	var existingItem domain.CartItem
	err = r.db.
		Where("cart_id = ? AND product_id = ?", cart.ID, cartItem.ProductID).
		First(&existingItem).Error

	if err == nil {
		existingItem.Quantity += cartItem.Quantity

		if err := r.db.Save(&existingItem).Error; err != nil {
			return nil, err
		}

		return &existingItem, nil
	}

	cartItem.CartID = cart.ID

	if err := r.db.Create(cartItem).Error; err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (r *CartRepository) ClearCart(userId int64) error {
	return r.db.Delete(&domain.Cart{}, "user_id = ?", userId).Error
}
