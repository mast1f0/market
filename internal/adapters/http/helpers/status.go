package helpers

import (
	"errors"
	"net/http"

	"market/internal/core/ports"
	"market/internal/core/service"
)

func StatusFromError(err error) int {
	switch {
	case errors.Is(err, service.ErrForbidden):
		return http.StatusForbidden

	case isNotFound(err):
		return http.StatusNotFound

	case isConflict(err):
		return http.StatusConflict

	case isBadRequest(err):
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}

func RespondErr(w http.ResponseWriter, err error) {
	RespondError(w, StatusFromError(err), err.Error())
}

func isNotFound(err error) bool {
	return errors.Is(err, service.ErrProductNotFound) ||
		errors.Is(err, service.ErrOrderNotFound) ||
		errors.Is(err, service.ErrCategoryNotFound) ||
		errors.Is(err, service.ErrCartNotFound) ||
		errors.Is(err, service.ErrCartItemNotFound) ||
		errors.Is(err, ports.ErrNotFound) ||
		errors.Is(err, ports.ErrOrderNotFound) ||
		errors.Is(err, ports.ErrCategoryNotFound) ||
		errors.Is(err, ports.ErrCartNotFound) ||
		errors.Is(err, ports.ErrCartItemNotFound)
}

func isConflict(err error) bool {
	return errors.Is(err, service.ErrCategoryExists) ||
		errors.Is(err, ports.ErrAlreadyExists) ||
		errors.Is(err, ports.ErrConflict) ||
		errors.Is(err, ports.ErrCategoryExists) ||
		errors.Is(err, ports.ErrCategoryAlreadyExists)
}

func isBadRequest(err error) bool {
	return errors.Is(err, service.ErrInvalidProductID) ||
		errors.Is(err, service.ErrInvalidOrderId) ||
		errors.Is(err, service.ErrInvalidUserID) ||
		errors.Is(err, service.ErrEmptyCart) ||
		errors.Is(err, service.ErrInvalidCategoryID) ||
		errors.Is(err, service.ErrInvalidCategoryName) ||
		errors.Is(err, service.ErrInvalidProduct) ||
		errors.Is(err, service.ErrInvalidItem) ||
		errors.Is(err, service.ErrInvalidQuantity) ||
		errors.Is(err, service.ErrInvalidCartID) ||
		errors.Is(err, service.ErrInvalidData) ||
		errors.Is(err, ports.ErrInvalidData)
}
