package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IGenreController interface {
	CreateGenre(c echo.Context) error
	UpdateGenre(c echo.Context) error
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

// ジャンル更新
func (gc *genreController) UpdateGenre(c echo.Context) error {
	id := c.Param("id")
	genreIdInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid genre id")
	}

	if genreIdInt < 0 {
		return c.JSON(http.StatusBadRequest, "genre id must be a positive integer")
	}

	genreIdUint := uint(genreIdInt)

	var genre model.Genre
	if err := c.Bind(&genre); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	genreRes, err := gc.gu.UpdateGenre(genre, genreIdUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, genreRes)
}
