package user

import (
	"testing"
)

func TestUserNew(t *testing.T) {
	emlStr := "test@example.com"
	passPlain := "Password@123"

	_, err := New(emlStr, passPlain)
	if err != nil {
		t.Errorf("error at creating New User %s", err)
	}
}
