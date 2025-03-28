package usecase

import (
	"backend/model"
	"backend/repository"
	"log"
)

type ITWordBookCase interface {
	GetAllWordBook() ([]model.WordBook, error)
	CreateWordBook(model.WordBook) (model.WordBookResponse, error)
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

// 英単語帳を新規作成
func (wu *wordBookUseCase) CreateWordBook(wordBook model.WordBook) (model.WordBookResponse, error) {
	if err := wu.wr.CreateWordBook(&wordBook); err != nil {
		log.Printf("failed to create word book: %v", err)
		return model.WordBookResponse{}, err
	}

	resWordBook := model.WordBookResponse{
		ID:           wordBook.ID,
		Title:        wordBook.Title,
		Description:  wordBook.Description,
		AdminId:      wordBook.AdminId,
		GenreId:      wordBook.GenreId,
		DifficultyId: wordBook.DifficultyId,
		CreatedAt:    wordBook.CreatedAt,
		UpdatedAt:    wordBook.UpdatedAt,
	}

	return resWordBook, nil
}