package user

import "github.com/yamaga-shu/jwt-go-next/api-service/domain/entity/user"

type repository interface {
	Store(*user.User) error
}
