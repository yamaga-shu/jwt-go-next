package user

import "github.com/yamaga-shu/jwt-go-next/api-service/domain/entity/user"

// Service is service for User
type Service struct {
	repository repository
}

// Store stores given user to Database
func (s Service) Store(u *user.User) error {
	return s.repository.Store(u)
}
