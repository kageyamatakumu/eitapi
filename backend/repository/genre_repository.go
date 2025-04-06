package repository

import (
	"backend/model"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IGenreRepository interface {
	CreateGenre(genre *model.Genre) error
	UpdateGenre(genre *model.Genre, genreId uint) error
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

// ジャンルを更新
func (gr *genreRepository) UpdateGenre(genre *model.Genre, genreId uint) error {
	result := gr.db.Model(genre).Clauses(clause.Returning{}).Where("id = ?", genreId).Updates(
		model.Genre{
			GenreName: genre.GenreName,
		},
	)

	if result.Error != nil {
		log.Printf("failed to update genre with id %d: %v", genreId, result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		log.Printf("genre id %d not found\n", genreId)
		return gorm.ErrRecordNotFound
	}

	log.Printf("successfully updated genre with id %d\n", genreId)

	return nil
}
