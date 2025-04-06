package usecase

import (
	"backend/model"
	"backend/repository"
	"backend/validator"
)

type ITWordBookUsecase interface {
	GetAllWordBook() ([]model.WordBook, error)
	CreateWordBook(model.WordBook) (model.WordBookResponse, error)
	DeleteWordBook(wordBookId uint) error
}

type wordBookUsecase struct {
	wr  repository.IWordBookRepository
	wbv validator.IWordBookValidator
}

func NewWordBookUsecase(wr repository.IWordBookRepository, wbv validator.IWordBookValidator) ITWordBookUsecase {
	return &wordBookUsecase{wr, wbv}
}

// 英単語帳を全て取得
func (wu *wordBookUsecase) GetAllWordBook() ([]model.WordBook, error) {
	var wordBooks []model.WordBook
	if err := wu.wr.GetAllWordBook(&wordBooks); err != nil {
		return nil, err
	}
	return wordBooks, nil
}

// 英単語帳を新規作成
func (wu *wordBookUsecase) CreateWordBook(wordBook model.WordBook) (model.WordBookResponse, error) {
	if err := wu.wbv.ValidateTitle(wordBook.Title); err != nil {
		return model.WordBookResponse{}, err
	}

	if err := wu.wr.CreateWordBook(&wordBook); err != nil {
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

// 英単語帳を削除
func (wu *wordBookUsecase) DeleteWordBook(wordBookId uint) error {
	if err := wu.wr.DeleteWordBook(wordBookId); err != nil {
		return err
	}

	return nil
}
