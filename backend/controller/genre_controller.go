package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IGenreController interface {
	CreateGenre(c echo.Context) error
}

type genreController struct {
	gu usecase.IGenreUsecase
}

func NewGenreController(gu usecase.IGenreUsecase) IGenreController {
	return &genreController{gu}
}

// ジャンル作成
func (gc *genreController) CreateGenre(c echo.Context) error {
	var genre model.Genre

	if err := c.Bind(&genre); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	genreRes, err := gc.gu.CreateGenre(genre)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, genreRes)
}
