package user

import (
	"context"

	"github.com/yamaga-shu/jwt-go-next/api-service/domain/entity/user"
)

type IUserRepository interface {
	Store(ctx context.Context, u *user.User) error
}
