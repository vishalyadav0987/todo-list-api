package auth

import "time"

type User struct {
	id           string
	name         string
	email        string
	passwordHash string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewUser(
	id string,
	name Name,
	email Email,
	passwordHash string,
) *User {

	now := time.Now()
	return &User{
		id:           id,
		name:         string(name),
		email:        string(email),
		passwordHash: passwordHash,
		createdAt:    now,
		updatedAt:    now,
	}
}

// ✅ Getters (Read-Only Access)
func (u *User) ID() string {
	return u.id
}

func (u *User) Name() Name {
	return Name(u.name)
}

func (u *User) Email() Email {
	return Email(u.email)
}

func (u *User) ChangePassword(newHash string) {
	u.passwordHash = newHash
	u.updatedAt = time.Now()
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
