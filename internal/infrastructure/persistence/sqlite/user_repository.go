package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vishalyadav0987/todo-list-api/internal/domain/auth"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// ✅ Implement Save
// implements to UserRepo interface
func (r *UserRepository) Save(ctx context.Context, user *auth.User) error {
	query := `INSERT INTO users (id, name, email, password, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query, user.ID(), user.Name(), user.Email(), user.PasswordHash(), user.CreatedAt(), user.UpdatedAt())
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email auth.Email) (*auth.User, error) {
	query := `SELECT id, name, email, password, createdAt, updatedAt FROM users WHERE email = $1`

	row := r.db.QueryRowContext(ctx, query, string(email))

	var (
		id        string
		name      string
		emailStr  string
		password  string
		createdAt string
		updatedAt string
	)

	err := row.Scan(
		&id,
		&name,
		&emailStr,
		&password,
		&createdAt,
		&updatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, auth.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	emailValidate, _ := auth.NewEmail(emailStr)
	nameValidate, _ := auth.NewName(name)

	user := auth.NewUser(
		id, nameValidate, emailValidate, password,
	)

	return user, nil
}

func (r *UserRepository) FindById(ctx context.Context, userId string) (*auth.User, error) {
	query := `SELECT id, name, email, password, createdAt, updatedAt FROM users WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, userId)

	var (
		id        string
		name      string
		emailStr  string
		password  string
		createdAt string
		updatedAt string
	)

	err := row.Scan(
		&id,
		&name,
		&emailStr,
		&password,
		&createdAt,
		&updatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, auth.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	emailValidate, _ := auth.NewEmail(emailStr)
	nameValidate, _ := auth.NewName(name)

	user := auth.NewUser(
		id, nameValidate, emailValidate, password,
	)

	return user, nil
}
