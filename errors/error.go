package errors

import "errors"

var (
	ErrInternal = errors.New("unknown error")
	ErrNotFound = errors.New("not found")
)
