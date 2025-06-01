package user

import (
	"testing"
)

func TestUserNew(t *testing.T) {
	emailStr := "test@example.com"
	passPlain := "Password@123"

	_, err := New(emailStr, passPlain)
	if err != nil {
		t.Errorf("error at creating New User %s", err)
	}
}
