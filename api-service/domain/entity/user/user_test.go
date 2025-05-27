package user

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	pass := "sample password"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		t.Fatal(err)
	}

	if !checkPasswordHash(pass, hashedPass) {
		t.Error("ckeckPasswordHash doesn't work. should be true, but got false")
	}

	if checkPasswordHash("incorrect password", hashedPass) {
		t.Error("ckeckPasswordHash doesn't work. should be false, but got true")
	}
}

func TestUserNew(t *testing.T) {
	emailStr := "test@example.com"
	passPlain := "password"

	u, err := New(emailStr, passPlain)
	if err != nil {
		t.Errorf("error at creating New User %s", err)
	}

	if u.email.String() != emailStr {
		t.Errorf("expected User.email.String() == emailStr, but doesn't: %s", u.email.String())
	}
}
