package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTodoNotFound       = errors.New("todo not found")
	ErrInvalidTodoId      = errors.New("invalid todo id")
)
