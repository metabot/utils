package utils

import "errors"

var (
	ErrDuplicate = errors.New("duplicate id")
	ErrNotFound = errors.New("not found")
	ErrInvalidRequest = errors.New("failed validation")
)
