package valobj

import (
	"testing"
)

func TestEmailNew(t *testing.T) {
	emlStr := "test@valid-email.com"

	eml, err := NewEmail(emlStr)
	if err != nil {
		t.Fatal(err)
	}

	if eml.value != emlStr {
		t.Errorf("Email.value must be same, but passed %s", eml.value)
	}
}
