package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/yamaga-shu/jwt-go-next/api-service/domain/entity/user"
)

// Service: User のドメインサービスを表現する構造体
type Service struct {
	repo IUserRepository
}

// NewDomainService: User のドメインサービスのコンストラクタ
func NewDomainService(UserRepo IUserRepository) Service {
	return Service{
		repo: UserRepo,
	}
}

// Store: ユーザーをDBに保存するサービスメソッド
func (s Service) Store(ctx context.Context, u *user.User) error {
	// UserId の設定確認
	if u.UserId == uuid.Nil {
		return errors.New("user.Service.Store: UserId must be a valid UUID")
	}

	// リポジトリ経由でユーザーを登録
	return s.repo.Store(ctx, u)
}
