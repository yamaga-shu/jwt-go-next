package valobj

import (
	"errors"
	"regexp"
)

var InvalidEmailErr = errors.New("invalid email format")
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Email: メールアドレスを表現する構造体
type Email struct {
	value string
}

// NewEmail は、形式が正しい場合のみ Email を生成します
func NewEmail(value string) (*Email, error) {
	if !emailRegex.MatchString(value) {
		return nil, InvalidEmailErr
	}
	return &Email{value: value}, nil
}

func (e *Email) String() string {
	return e.value
}
