package user

import (
	"github.com/google/uuid"
	"github.com/yamaga-shu/jwt-go-next/api-service/domain/valobj/email"
	"github.com/yamaga-shu/jwt-go-next/api-service/domain/valobj/password"
)

// User: 一般ユーザー
type User struct {
	userId     uuid.UUID
	email      email.Email
	hashedPass string
}

func New(strEmail, plainPass string) (*User, error) {
	// id の生成
	id := uuid.New()

	// エラーの初期化
	var err error

	// パスワードの検証
	var pass *password.Password
	pass, err = password.New(plainPass)
	if err != nil {
		return nil, err
	}
	// # パスワードのハッシュ化
	var hashedPass string
	hashedPass, err = pass.Hash()
	if err != nil {
		return nil, err
	}

	// メールの検証
	var userEmail *email.Email
	userEmail, err = email.New(strEmail)
	if err != nil {
		return nil, err
	}

	return &User{
		userId:     id,
		email:      *userEmail,
		hashedPass: hashedPass,
	}, nil
}
