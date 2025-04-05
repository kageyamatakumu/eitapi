package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IDifficultyController interface {
	CreateDifficulty(c echo.Context) error
	UpdateDifficulty(c echo.Context) error
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

	resDifficulty, err := dc.du.CreateDifficulty(difficulty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resDifficulty)
}

// 難易度を更新
func (dc *difficultyController) UpdateDifficulty(c echo.Context) error {
	id := c.Param("id")
	difficultyIdInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid difficulty id")
	}

	if difficultyIdInt < 0 {
		return c.JSON(http.StatusBadRequest, "difficulty id must be a positive integer")
	}

	difficultyIdUint := uint(difficultyIdInt)

	difficulty := model.Difficulty{}
	if err := c.Bind(&difficulty); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resDifficulty, err := dc.du.UpdateDifficulty(difficulty, difficultyIdUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resDifficulty)
}
