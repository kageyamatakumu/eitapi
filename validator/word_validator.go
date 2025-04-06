package validator

import (
	"fmt"
	"regexp"
)

type IWordValidator interface {
	ValidateEnglishWord(englishWord string) error
	ValidateJapaneseTranslation(japaneseTranslation string) error
	ValidatePronunciation(pronunciation string) error
}

type wordValidator struct{}

func NewWordValidator() IWordValidator {
	return &wordValidator{}
}

// 英単語
func (wv *wordValidator) ValidateEnglishWord(englishWord string) error {
	if englishWord == "" {
		return fmt.Errorf("英単語を入力してくだい")
	}

	if len(englishWord) > 255 {
		return fmt.Errorf("英単語が長すぎます")
	}

	englishWordRegex := `^[a-zA-Z]+(?:[-'][a-zA-Z]+)*$`
	re, err := regexp.Compile(englishWordRegex)
	if err != nil {
		return fmt.Errorf("内部エラー: 英単語の形式チェックに失敗しました")
	}
	if !re.MatchString(englishWord) {
		return fmt.Errorf("英単語はアルファベットのみで入力してください")
	}

	return nil
}

// 和訳
func (wv *wordValidator) ValidateJapaneseTranslation(japaneseTranslation string) error {
	if japaneseTranslation == "" {
		return fmt.Errorf("和訳を入力してくだい")
	}

	if len(japaneseTranslation) > 255 {
		return fmt.Errorf("和訳が長すぎます")
	}

	return nil
}

// 発音
func (wv *wordValidator) ValidatePronunciation(pronunciation string) error {
	if pronunciation == "" {
		return fmt.Errorf("発音を入力してくだい")
	}

	if len(pronunciation) > 255 {
		return fmt.Errorf("発音が長すぎます")
	}

	return nil
}
