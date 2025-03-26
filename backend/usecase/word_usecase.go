package usecase

import (
	"backend/model"
	"backend/repository"
	"log"
)

type IWordCase interface {
	GetAllWordsForWordBook(wordBookID uint) ([]model.Word, error)
}

type wordUseCase struct {
	wr repository.IWordRepository
}

func NewWordUsecase(wr repository.IWordRepository) IWordCase {
	return &wordUseCase{wr}
}

// 英単語帳に紐づく英単語を全て取得
func (wu *wordUseCase) GetAllWordsForWordBook(wordBookID uint) ([]model.Word, error) {
	var words []model.Word
	if err := wu.wr.GetAllWordsForWordBook(wordBookID, &words); err != nil {
		log.Printf("failed to get all words for word book: %v", err)
		return nil, err
	}

	return words, nil
}
