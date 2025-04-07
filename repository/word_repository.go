package repository

import (
	"fmt"
	"log"

	"backend/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IWordRepository interface {
	GetAllWordsForWordBook(wordBookId uint, words *[]model.Word) error
	GetWordById(wordId uint, word *model.Word) error
	CreateMultipleWords(words *[]model.Word) error
	UpdateWord(word *model.Word, wordId uint) error
	DeleteWord(wordId uint) error
}

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) IWordRepository {
	return &wordRepository{db}
}

// 英単語帳に紐づく英単語を全て取得
func (wr *wordRepository) GetAllWordsForWordBook(wordBookId uint, words *[]model.Word) error {
	if err := wr.db.Model(&model.Word{}).Where("word_book_id = ?", wordBookId).Find(words).Error; err != nil {
		log.Printf("failed to get all words for word book with id %d: %v", wordBookId, err)
		return err
	}
	return nil
}

// 英単語を取得
func (wr *wordRepository) GetWordById(wordId uint, word *model.Word) error {
	if err := wr.db.Model(model.Word{}).Where("id = ?", wordId).First(word).Error; err != nil {
		log.Printf("failed to get word with id %d: %v", wordId, err)
		return err
	}

	return nil
}

// 英単語を新規作成
func (wr *wordRepository) CreateMultipleWords(words *[]model.Word) error {
	if err := wr.db.Model(model.Word{}).Create(words).Error; err != nil {
		log.Printf("failed to create words: %v\n", err)
		return err
	}

	log.Printf("successfully created word\n")

	return nil
}

// 英単語を更新
func (wr *wordRepository) UpdateWord(word *model.Word, wordId uint) error {
	result := wr.db.Model(word).Clauses(clause.Returning{}).Where("id = ?", wordId).Updates(
		model.Word{
			EnglishWord:         word.EnglishWord,
			JapaneseTranslation: word.JapaneseTranslation,
			Pronunciation:       word.Pronunciation,
			ExampleSentence:     word.ExampleSentence,
		})

	if result.Error != nil {
		log.Printf("failed to update word with id %d: %v\n", wordId, result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		log.Printf("failed to update word with id %d: %v\n", wordId, result.Error)
		return fmt.Errorf("object does not exist")
	}

	log.Printf("successfully updated word with id %d\n", wordId)

	return nil
}

// 英単語を削除
func (wr *wordRepository) DeleteWord(wordId uint) error {
	result := wr.db.Model(model.Word{}).Where("id = ?", wordId).Delete(&model.Word{})
	if result.Error != nil {
		log.Printf("failed to delete word with id %d: %v\n", wordId, result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		log.Printf("word with id %d not found\n", wordId)
		return fmt.Errorf("object does not exist")
	}

	log.Printf("successfully deleted word with id %d\n", wordId)

	return nil
}
