package validator

import (
	"fmt"
	"regexp"
	"unicode"
)

type IAdminValidator interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
}

type adminValidator struct{}

func NewAdminValidator() IAdminValidator {
	return &adminValidator{}
}

// メールアドレス
func (av *adminValidator)ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("メールアドレスを入力してください")
	}

	if len(email) > 255 {
		return fmt.Errorf("メールアドレスが長すぎます")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re, err := regexp.Compile(emailRegex)
	if err != nil {
		return fmt.Errorf("内部エラー: メールアドレスの形式チェックに失敗しました")
	}
	if !re.MatchString(email) {
		return fmt.Errorf("メールアドレスの形式が不正です")
	}

	return nil
}

// パスワード
func (av *adminValidator)ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("パスワードを入力してください")
	}

	if len(password) < 8 || len(password) > 16 {
		return fmt.Errorf("パスワードは8文字以上16文字以下である必要があります")
	}

	var hasLetter, hasDigit, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsLetter(ch):
			hasLetter = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasLetter || !hasDigit || !hasSpecial {
		return fmt.Errorf("パスワードは、英字、数字、記号をそれぞれ1文字以上含む必要があります")
	}

	return nil
}
