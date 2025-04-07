package validator

import (
	"backend/model"
	"fmt"
)

type IDifficultyValidator interface {
	ValidateDifficultyLevel(difficultyLevel model.DifficultyLevel) error
}

type difficultyValidator struct{}

func NewDifficultyValidator() IDifficultyValidator {
	return &difficultyValidator{}
}

// 難易度の設定範囲
func (dv *difficultyValidator) ValidateDifficultyLevel(difficultyLevel model.DifficultyLevel) error {
	if difficultyLevel < model.Easy || difficultyLevel > model.Hard {
		return fmt.Errorf("難易度は1から3の範囲で指定してください")
	}

	return nil
}
