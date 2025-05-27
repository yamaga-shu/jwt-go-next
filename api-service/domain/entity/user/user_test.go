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
