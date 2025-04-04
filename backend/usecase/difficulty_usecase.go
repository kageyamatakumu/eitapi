package usecase

import (
	"backend/model"
	"backend/repository"
)

type IDifficultyUsecase interface {
	CreateDifficulty(difficulty model.Difficulty) (model.DifficultyResponse, error)
}

type difficultyUsecase struct {
	dr repository.IDifficultyRepository
}

func NewDifficultyUsecase(dr repository.IDifficultyRepository) IDifficultyUsecase {
	return &difficultyUsecase{dr}
}

// 難易度を作成
func (du *difficultyUsecase) CreateDifficulty(difficulty model.Difficulty) (model.DifficultyResponse, error) {
	if err := du.dr.CreateDifficulty(&difficulty); err != nil {
		return model.DifficultyResponse{}, err
	}

	resIDifficulty := model.DifficultyResponse{
		ID:              difficulty.ID,
		DifficultyLevel: difficulty.DifficultyLevel,
	}

	return resIDifficulty, nil
}
