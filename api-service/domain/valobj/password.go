package valobj

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var InvalidPasswordErr = errors.New("invalid password format")

// パスワードの正規表現
// 次の条件を満たす
var (
	hankakuRegex  = regexp.MustCompile(`^[\x01-\x7E]+$`) // 半角文字で構成される
	notSpaceRegex = regexp.MustCompile(`^\S+$`)          // 空白文字を含まない(空白文字以外で構成される)
	numberRegex   = regexp.MustCompile(`\d`)             // 数字を含む
	lowerRegex    = regexp.MustCompile(`[a-z]`)          // 英小文字を含む
	upperRegex    = regexp.MustCompile(`[A-Z]`)          // 英大文字を含む
	symbolRegex   = regexp.MustCompile(`[@$!%*?&]`)      // 記号を含む
	lengthRegex   = regexp.MustCompile(`^.{6,24}$`)      // 6文字以上24文字以下で構成される
)

// Password: パスワードを表現する構造体
type Password struct {
	plain string
}

// NewPassword: パスワードの形式を確認し、パスワード構造体を生成する
func NewPassword(plain string) (*Password, error) {
	if !(hankakuRegex.MatchString(plain) &&
		notSpaceRegex.MatchString(plain) &&
		numberRegex.MatchString(plain) &&
		lowerRegex.MatchString(plain) &&
		upperRegex.MatchString(plain) &&
		symbolRegex.MatchString(plain) &&
		lengthRegex.MatchString(plain)) {
		return nil, InvalidPasswordErr
	}

	return &Password{
		plain: plain,
	}, nil
}

// Plain: パスワードの平文を返却するメソッド
func (p Password) Plain() string {
	return p.plain
}

// Hash: 保持している平文のパスワードをハッシュ化して返却するメソッド
func (p Password) Hash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p.plain), 14)
	return string(bytes), err
}

// Check: ハッシュ化されたパスワードが、保持している平文のパスワードと一致しているか
// 確認するメソッド
func (p Password) Check(hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.plain))
	return err == nil
}
