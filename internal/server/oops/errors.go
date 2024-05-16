package oops

import "errors"

var (
	ErrUserExist       = errors.New("user already exist error")
	ErrUserNotFound    = errors.New("user not found error")
	ErrTokenNotFound   = errors.New("token not found error")
	ErrPasswordInvalid = errors.New("password invalid error")
)
