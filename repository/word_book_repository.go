package repository

import (
	"backend/model"
	"log"

	"gorm.io/gorm"
)

type IWordBookRepository interface {
	GetAllWordBook(wordBooks *[]model.WordBook) error
	CreateWordBook(wordBook *model.WordBook) error
	DeleteWordBook(wordBookId uint) error
}

type wordBookRepository struct {
	db *gorm.DB
}

func NewWordBookRepository(db *gorm.DB) IWordBookRepository {
	return &wordBookRepository{db}
}

// 英単語帳を全て取得
func (wr *wordBookRepository) GetAllWordBook(wordBooks *[]model.WordBook) error {
	if err := wr.db.Model(model.WordBook{}).Preload("Genre").Preload("Difficulty").Find(wordBooks).Error; err != nil {
		log.Printf("failed to get all word books: %v", err)
		return err
	}
	return nil
}

// 英単語帳を新規作成
func (wr *wordBookRepository) CreateWordBook(wordBook *model.WordBook) error {
	if err := wr.db.Model(model.WordBook{}).Create(wordBook).Error; err != nil {
		log.Printf("failed to create word book: %v\n", err)
		return err
	}

	log.Printf("successfully created word book\n")

	return nil
}

// 英単語帳を削除
func (wr *wordBookRepository) DeleteWordBook(wordBookId uint) error {
	result := wr.db.Model(model.WordBook{}).Where("id = ?", wordBookId).Delete(&model.WordBook{})
	if result.Error != nil {
		log.Printf("failed to delete word book with id %d: %v\n", wordBookId, result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		log.Printf("word book id %d not found\n", wordBookId)
		return gorm.ErrRecordNotFound
	}

	log.Printf("successfully deleted word book with id %d\n", wordBookId)

	return nil
}
