package authapp

import (
	"context"

	"github.com/vishalyadav0987/todo-list-api/internal/domain/auth"
)

type PasswordHasher interface {
	GenerateHash(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

type IDGenerator interface {
	Generate() string
}

type RegisterUsecase struct {
	repo   auth.UserRepository
	hasher PasswordHasher
	idGen  IDGenerator
}

func NewRegisterUsecase(
	repo auth.UserRepository,
	hasher PasswordHasher,
	idGen IDGenerator,
) *RegisterUsecase {
	return &RegisterUsecase{
		repo:   repo,
		hasher: hasher,
		idGen:  idGen,
	}
}

type RegisterRequest struct {
	Email    string
	Password string
	Name     string
}

func (uc *RegisterUsecase) Execute(
	ctx context.Context,
	req RegisterRequest,
) (*auth.User, error) {

	name, err := auth.NewName(req.Name)
	if err != nil {
		return nil, err
	}

	email, err := auth.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	existingUser, _ := uc.repo.FindByEmail(ctx, auth.Email(req.Email))
	if existingUser != nil {
		return nil, auth.ErrUserAlreadyExists
	}

	hashedPassword, err := uc.hasher.GenerateHash(req.Password)
	if err != nil {
		return nil, err
	}

	id := uc.idGen.Generate()

	user := auth.NewUser(
		id,
		name, email, hashedPassword,
	)

	err = uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil

}
