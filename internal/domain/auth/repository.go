package auth

import "context"

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email Email) (*User, error)
	FindById(ctx context.Context, userId string) (*User, error)
}
