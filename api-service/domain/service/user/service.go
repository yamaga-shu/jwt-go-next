package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yamaga-shu/jwt-go-next/api-service/domain/entity/user"
)

// Service: User のドメインサービスを表現する構造体
type Service struct {
	repository repository
}

// Store: ユーザーをDBに保存するサービスメソッド
func (s Service) Store(u *user.User) error {
	// UserId の設定確認
	if u.UserId == uuid.Nil {
		return errors.New("user.Service.Store: UserId must be a valid UUID")
	}

	// パスワードのハッシュ化
	hashed, err := u.Password.Hash()
	if err != nil {
		return err
	}

	// リポジトリ経由でユーザーを登録
	return s.repository.Store(u.UserId, u.Email.String(), hashed)
}
