package service

import "errors"

var (
	ErrAlreadyExists       = errors.New("already exists")
	ErrConflict            = errors.New("conflict")
	ErrInvalidData         = errors.New("invalid data")
	ErrFailedToLoadProduct = errors.New("failed to load product")

	ErrFailedToLoadCart       = errors.New("failed to load cart")
	ErrFailedToSaveCart       = errors.New("failed to save cart")
	ErrFailedToDeleteCartItem = errors.New("failed to delete cart item")
	ErrFailedToUpdateCartItem = errors.New("failed to update cart item")
	ErrFailedToClearCart      = errors.New("failed to clear cart")
	ErrFailedToLoadCartItem   = errors.New("failed to load cart item")
)
