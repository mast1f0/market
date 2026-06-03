package mock

import "market/internal/core/domain"

type CartRepoMock struct {
	DeleteCartItemById func(userId int64, productId int64) error
	AddCartItemA       func(userId int64, item *domain.CartItem) (*domain.CartItem, error)
	CreateCartByUserId func(userId int64) (*domain.Cart, error)
}

func (c *CartRepoMock) CreateCart(id int64) (*domain.Cart, error) {
	return c.CreateCartByUserId(id)
}

func (c *CartRepoMock) GetCartWithItems(userID int64) (*domain.Cart, error) {
	return nil, nil
}

func (c *CartRepoMock) ClearCart(userId int64) error {
	return nil
}

func (c *CartRepoMock) FindCartItem(cartID, productID int64) (*domain.CartItem, error) {
	return nil, nil
}

func (r *CartRepoMock) DeleteCartItem(userId int64, productId int64) error {
	return r.DeleteCartItemById(userId, productId)
}

func (r *CartRepoMock) UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error) {
	return nil, nil
}

func (r *CartRepoMock) AddCartItem(userID int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	return r.AddCartItemA(userID, cartItem)
}
