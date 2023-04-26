package utils

import (
	"errors"
)

type CustomErrorWrapper struct {
	Message string `json:"message"` // Human readable message for clients
	Code    int    `json:"-"`       // HTTP Status code. We use `-` to skip json marshaling.
	Err     error  `json:"-"`       // The original error. Same reason as above.
}

// Returns Message if Err is nil
func (err CustomErrorWrapper) Error() string {
	if err.Code == 707 {
		return err.Message
	}
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.Message
}

func (err CustomErrorWrapper) GetCode() int {
	if err.Code != 0 {
		return 500
	}
	return err.Code
}

func (err CustomErrorWrapper) Unwrap() error {
	return err.Err // Returns inner error
}

// Returns the inner most CustomErrorWrapper
func (err CustomErrorWrapper) Dig() CustomErrorWrapper {
	var ew CustomErrorWrapper
	if errors.As(err.Err, &ew) {
		// Recursively digs until wrapper error is not CustomErrorWrapper
		return ew.Dig()
	}
	return err
}

func NewErrorWrapper(code int, err error, message string) error {
	return CustomErrorWrapper{
		Message: message,
		Code:    code,
		Err:     err,
	}
}
