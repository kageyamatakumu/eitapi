package repository

import (
	"backend/model"
	"log"

	"gorm.io/gorm"
)

type IGenreRepository interface {
	CreateGenre(genre *model.Genre) error
}

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) IGenreRepository {
	return &genreRepository{db}
}

// ジャンルを作成
func (gr *genreRepository) CreateGenre(genre *model.Genre) error {
	if err := gr.db.Model(&model.Genre{}).Create(genre).Error; err != nil {
		log.Printf("failed to create genre: %v\n", err)
		return err
	}

	log.Printf("successfully created genre\n")

	return nil
}
