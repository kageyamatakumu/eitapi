package usecase

import (
	"backend/model"
	"backend/repository"
	"log"
)

type IWordCase interface {
	GetAllWordsForWordBook(wordBookID uint) ([]model.Word, error)
	CreateWord(word model.Word) (model.WordRes, error)
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

// 英単語を新規作成
func (wu *wordUseCase) CreateWord(word model.Word) (model.WordRes, error) {
	if err := wu.wr.CreateWord(&word); err != nil {
		log.Printf("failed to create word: %v", err)
		return model.WordRes{}, err
	}

	resWord := model.WordRes{
		ID:                  word.ID,
		EnglishWord:         word.EnglishWord,
		JapaneseTranslation: word.JapaneseTranslation,
		Pronunciation:       word.Pronunciation,
		ExampleSentence:     word.ExampleSentence,
		WordBookId:          word.WordBookId,
	}

	return resWord, nil
}
