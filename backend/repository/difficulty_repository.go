package repository

import (
	"backend/model"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDifficultyRepository interface {
	GetAllDifficulties(difficulties *[]model.Difficulty) error
	CreateDifficulty(difficulty *model.Difficulty) error
	UpdateDifficulty(difficulty *model.Difficulty, difficultyId uint) error
	DeleteDifficulty(difficultyId uint) error
}

type difficultyRepository struct {
	db *gorm.DB
}

func NewDifficultyRepository(db *gorm.DB) IDifficultyRepository {
	return &difficultyRepository{db}
}

// 難易度を全て取得
func (dr *difficultyRepository) GetAllDifficulties(difficulties *[]model.Difficulty) error {
	if err := dr.db.Model(&model.Difficulty{}).Find(difficulties).Error; err != nil {
		log.Printf("failed to get all difficulties: %v", err)
		return err
	}

	return nil
}

// 難易度を作成
func (dr *difficultyRepository) CreateDifficulty(difficulty *model.Difficulty) error {
	if err := dr.db.Create(difficulty).Error; err != nil {
		log.Printf("failed to create difficulty: %v\n", err)
		return err
	}

	log.Printf("successfully created difficulty\n")

	return nil
}

// 難易度を更新
func (dr *difficultyRepository) UpdateDifficulty(difficulty *model.Difficulty, difficultyId uint) error {
	result := dr.db.Model(difficulty).Clauses(clause.Returning{}).Where("id = ?", difficultyId).Updates(
		model.Difficulty{
			DifficultyLevel: difficulty.DifficultyLevel,
		})

	if result.Error != nil {
		log.Printf("failed to update difficulty with id %d: %v", difficultyId, result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		log.Printf("no rows affected when updating difficulty with id %d", difficultyId)
		return gorm.ErrRecordNotFound
	}

	log.Printf("successfully updated difficulty with id %d\n", difficultyId)

	return nil
}

// 難易度を削除
func (dr *difficultyRepository) DeleteDifficulty(difficultyId uint) error {
	result := dr.db.Where("id = ?", difficultyId).Delete(&model.Difficulty{})
	if result.Error != nil {
		log.Printf("failed to delete difficulty with id %d: %v", difficultyId, result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		log.Printf("difficulty with id %d not found", difficultyId)
		return gorm.ErrRecordNotFound
	}

	log.Printf("successfully deleted difficulty with id %d\n", difficultyId)

	return nil
}