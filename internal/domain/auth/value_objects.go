package auth

import (
	"errors"
	"regexp"
)

type Email string
type Name string

func NewName(value string) (Name, error) {
	if value == "" {
		return "", errors.New("name required")
	}

	if len(value) < 3 {
		return Name(value), errors.New("name should greter than 3")
	}

	return Name(value), nil
}

func NewEmail(value string) (Email, error) {
	if value == "" {
		return "", errors.New("email required")
	}

	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	matched, _ := regexp.MatchString(regex, value)

	if !matched {
		return "", errors.New("invalid email format")
	}

	return Email(value), nil
}
