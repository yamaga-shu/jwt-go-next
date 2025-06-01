package user

import (
	"github.com/google/uuid"
)

type repository interface {
	Store(id uuid.UUID, email string, hashedPass string) error
}
