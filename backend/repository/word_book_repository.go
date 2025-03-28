package repository

import (
	"backend/model"
	"log"

	"gorm.io/gorm"
)

type IWordBookRepository interface {
	GetAllWordBook(wordBooks *[]model.WordBook) error
	CreateWordBook(wordBook *model.WordBook) error
}

type wordBookRepository struct {
	db *gorm.DB
}

func NewWordBookRepository(db *gorm.DB) IWordBookRepository {
	return &wordBookRepository{db}
}

// 英単語帳を全て取得
func (wr *wordBookRepository) GetAllWordBook(wordBooks *[]model.WordBook) error {
	if err := wr.db.Find(wordBooks).Error; err != nil {
		log.Printf("failed to get all word books: %v", err)
		return err
	}
	return nil
}

// 英単語帳を新規作成
func (wr *wordBookRepository) CreateWordBook(wordBook *model.WordBook) error {
	if err := wr.db.Create(wordBook).Error; err != nil {
		log.Printf("failed to create word book: %v", err)
		return err
	}

	return nil
}
