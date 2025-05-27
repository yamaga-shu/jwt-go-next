package email

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type Email struct {
	value string
}

// NewEmail は、形式が正しい場合のみ Email を生成します
func (e *Email) New(value string) (Email, error) {
	if !emailRegex.MatchString(value) {
		return Email{}, errors.New("invalid email format")
	}
	return Email{value: value}, nil
}

func (e *Email) String() string {
	return e.value
}
