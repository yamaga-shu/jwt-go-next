package user

import "golang.org/x/crypto/bcrypt"

func hashPassword(plainPass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPass), 14)
	return string(bytes), err
}

func checkPasswordHash(plainPass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
	return err == nil
}
