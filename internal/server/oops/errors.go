package oops

import "errors"

var (
	ErrUserExist     = errors.New("user already exist error")
	ErrTokenNotFound = errors.New("token not found error")
)
