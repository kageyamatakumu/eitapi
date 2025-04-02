package validator

import "fmt"

type IWordBookValidator interface {
	ValidateTitle(title string) error
}

type wordBookValidator struct{}

func NewWordBookValidator() IWordBookValidator {
	return &wordBookValidator{}
}

// タイトル
func (wbv *wordBookValidator) ValidateTitle(title string) error {
	if title == "" {
		return fmt.Errorf("タイトルを入力してださい")
	}

	if len(title) > 255 {
		return fmt.Errorf("タイトルが長すぎきます")
	}

	return nil
}
