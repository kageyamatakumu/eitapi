package repository

import (
	"log"

	"backend/model"

	"gorm.io/gorm"
)

type IWordRepository interface {
	GetAllWordsForWordBook(wordBookID uint, words *[]model.Word) error
	CreateWord(word *model.Word) error
}

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) IWordRepository {
	return &wordRepository{db}
}

// 英単語帳に紐づく英単語を全て取得
func (wr *wordRepository) GetAllWordsForWordBook(wordBookID uint, words *[]model.Word) error {
	if err := wr.db.Model(&model.Word{}).Where("word_book_id = ?", wordBookID).Find(words).Error; err != nil {
		log.Printf("failed to get all words for word book: %v", err)
		return err
	}
	return nil
}

// 英単語を新規作成
func (wr *wordRepository) CreateWord(word *model.Word) error {
	if err := wr.db.Create(word).Error; err != nil {
		log.Printf("failed to create word: %v", err)
		return err
	}

	return nil
}
