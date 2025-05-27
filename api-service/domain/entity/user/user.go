package user

import (
	"github.com/google/uuid"
	"github.com/yamaga-shu/jwt-go-next/api-service/domain/valobj/email"
)

// User: 一般ユーザー
type User struct {
	userId     uuid.UUID
	email      email.Email
	hashedPass string
}

func (u User) New(strEmail, plainPass string) (*User, error) {
	// id の生成
	id := uuid.New()

	// パスワードのハッシュ化
	hashedPass, err := hashPassword(plainPass)
	if err != nil {
		return nil, err
	}

	// メールの検証
	email := email.Email{}
	email, err = email.New(strEmail)
	if err != nil {
		return nil, err
	}

	return &User{
		userId:     id,
		email:      email,
		hashedPass: hashedPass,
	}, nil
}
