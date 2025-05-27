package user

import (
	"github.com/google/uuid"
)

// User: 一般ユーザー
type User struct {
	userId     uuid.UUID
	email      string
	hashedPass string
}

func (u User) New(email, pass string) (*User, error) {
	id := uuid.New()

	hashedPass, err := hashPassword(pass)
	if err != nil {
		return nil, err
	}

	return &User{
		userId:     id,
		email:      email,
		hashedPass: hashedPass,
	}, nil
}
