package domains

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest            = errors.New("bad request parameters: %s")
	ErrUnauthorizedOperation = errors.New("unauthorized operation")
	ErrExistAlready          = errors.New("the request item has been exist already")

	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
)

func WrappErr(err error, msg string) error {
	if msg != "" {
		return fmt.Errorf("%w: "+msg, err)
	} else {
		return err
	}
}
