package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IDifficultyController interface {
	CreateDifficulty(c echo.Context) error
}

type difficultyController struct {
	du usecase.IDifficultyUsecase
}

func NewDifficultyController(du usecase.IDifficultyUsecase) IDifficultyController {
	return &difficultyController{du}
}

// 難易度を新規作成
func (dc *difficultyController) CreateDifficulty(c echo.Context) error {
	difficulty := model.Difficulty{}
	if err := c.Bind(&difficulty); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newDifficulty, err := dc.du.CreateDifficulty(difficulty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newDifficulty)
}
