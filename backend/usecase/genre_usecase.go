package usecase

import (
	"backend/model"
	"backend/repository"
)

type IGenreUsecase interface {
	CreateGenre(genre model.Genre) (model.GenreRes, error)
	UpdateGenre(genre model.Genre, genreId uint) (model.GenreRes, error)
}

type genreUsecase struct {
	gr repository.IGenreRepository
}

func NewGenreUsecase(gr repository.IGenreRepository) IGenreUsecase {
	return &genreUsecase{gr}
}

// ジャンルを作成
func (gu *genreUsecase) CreateGenre(genre model.Genre) (model.GenreRes, error) {
	if err := gu.gr.CreateGenre(&genre); err != nil {
		return model.GenreRes{}, err
	}

	genreRes := model.GenreRes{
		ID:        genre.ID,
		GenreName: genre.GenreName,
	}

	return genreRes, nil
}

// ジャンルを更新
func (gu *genreUsecase) UpdateGenre(genre model.Genre, genreId uint) (model.GenreRes, error) {
	if err := gu.gr.UpdateGenre(&genre, genreId); err != nil {
		return model.GenreRes{}, err
	}

	genreRes := model.GenreRes{
		ID:        genre.ID,
		GenreName: genre.GenreName,
	}

	return genreRes, nil
}
