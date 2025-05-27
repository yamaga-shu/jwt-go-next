package email

import (
	"testing"
)

func TestEmailNew(t *testing.T) {
	emailStr := "test@valid-email.com"

	email, err := New(emailStr)
	if err != nil {
		t.Fatal(err)
	}

	if email.value != emailStr {
		t.Errorf("Email.value must be same, but passed %s", email.value)
	}
}
