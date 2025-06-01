package user

import (
	"github.com/google/uuid"
	"github.com/yamaga-shu/jwt-go-next/api-service/domain/valobj"
)

// User: 一般ユーザー
type User struct {
	UserId   uuid.UUID
	Email    valobj.Email
	Password valobj.Password
}

// New: User構造体のコンストラクタ
func New(email, plainPass string) (*User, error) {
	// id の生成
	id := uuid.New()

	// エラーの初期化
	var err error

	// パスワードの検証
	var pass *valobj.Password
	pass, err = valobj.NewPassword(plainPass)
	if err != nil {
		return nil, err
	}

	// メールの検証
	var eml *valobj.Email
	eml, err = valobj.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		UserId:   id,
		Email:    *eml,
		Password: *pass,
	}, nil
}
