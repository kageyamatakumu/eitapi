package usecase

import (
	"backend/model"
	"backend/repository"
	"log"
)

type ITWordBookCase interface {
	GetAllWordBook() ([]model.WordBook, error)
}

type wordBookUseCase struct {
	wr repository.IWordBookRepository
}

func NewWordBookUsecase(wr repository.IWordBookRepository) ITWordBookCase {
	return &wordBookUseCase{wr}
}

// 英単語帳を全て取得
func (wu *wordBookUseCase) GetAllWordBook() ([]model.WordBook, error) {
	var wordBooks []model.WordBook
	if err := wu.wr.GetAllWordBook(&wordBooks); err != nil {
		log.Printf("failed to get all word books: %v", err)
		return nil, err
	}
	return wordBooks, nil
}
