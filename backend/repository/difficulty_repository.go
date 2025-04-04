package repository

import (
	"backend/model"
	"log"

	"gorm.io/gorm"
)

type IDifficultyRepository interface {
	CreateDifficulty(difficulty *model.Difficulty) error
}

type difficultyRepository struct {
	db *gorm.DB
}

func NewDifficultyRepository(db *gorm.DB) IDifficultyRepository {
	return &difficultyRepository{db}
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
