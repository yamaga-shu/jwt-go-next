package repository

import (
	"context"

	"github.com/yamaga-shu/jwt-go-next/api-service/domain/entity/user"
	"github.com/yamaga-shu/jwt-go-next/api-service/ent"
)

type UserRepository struct {
	db *ent.Client
}

func (ur UserRepository) Store(ctx context.Context, u *user.User) error {
	var err error
	var hashed string

	hashed, err = u.Password.Hash()
	if err != nil {
		return err
	}

	_, err = ur.db.User.
		Create().
		SetID(u.UserId).
		SetEmail(u.Email.String()).
		SetPassword(hashed).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}
