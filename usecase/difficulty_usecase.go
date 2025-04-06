package usecase

import (
	"backend/model"
	"backend/repository"
)

type IDifficultyUsecase interface {
	GetAllDifficulties()([]model.DifficultyResponse, error)
	CreateDifficulty(difficulty model.Difficulty) (model.DifficultyResponse, error)
	UpdateDifficulty(difficulty model.Difficulty, difficultyId uint) (model.DifficultyResponse, error)
	DeleteDifficulty(difficultyId uint) error
}

type difficultyUsecase struct {
	dr repository.IDifficultyRepository
}

func NewDifficultyUsecase(dr repository.IDifficultyRepository) IDifficultyUsecase {
	return &difficultyUsecase{dr}
}

// 難易度を全て取得
func(du *difficultyUsecase) GetAllDifficulties() ([]model.DifficultyResponse, error) {
	var difficulties []model.Difficulty
	if err := du.dr.GetAllDifficulties(&difficulties); err != nil {
		return nil, err
	}

	resDifficulties := make([]model.DifficultyResponse, len(difficulties))
	for i, difficulty := range difficulties {
		resDifficulties[i] = model.DifficultyResponse{
			ID:              difficulty.ID,
			DifficultyLevel: difficulty.DifficultyLevel,
		}
	}
	return resDifficulties, nil
}

// 難易度を作成
func (du *difficultyUsecase) CreateDifficulty(difficulty model.Difficulty) (model.DifficultyResponse, error) {
	if err := du.dr.CreateDifficulty(&difficulty); err != nil {
		return model.DifficultyResponse{}, err
	}

	resDifficulty := model.DifficultyResponse{
		ID:              difficulty.ID,
		DifficultyLevel: difficulty.DifficultyLevel,
	}

	return resDifficulty, nil
}

// 難易度を更新
func (du *difficultyUsecase) UpdateDifficulty(difficulty model.Difficulty, difficultyId uint) (model.DifficultyResponse, error) {
	if err := du.dr.UpdateDifficulty(&difficulty, difficultyId); err != nil {
		return model.DifficultyResponse{}, err
	}

	resDifficulty := model.DifficultyResponse{
		ID:              difficulty.ID,
		DifficultyLevel: difficulty.DifficultyLevel,
	}

	return resDifficulty, nil
}

// 難易度を削除
func (du *difficultyUsecase) DeleteDifficulty(difficultyId uint) error {
	if err := du.dr.DeleteDifficulty(difficultyId); err != nil {
		return err
	}

	return nil
}
