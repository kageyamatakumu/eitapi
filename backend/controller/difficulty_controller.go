package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IDifficultyController interface {
	GetAllDifficultiesPublic(c echo.Context) error
	GetAllDifficulties(c echo.Context) error
	CreateDifficulty(c echo.Context) error
	UpdateDifficulty(c echo.Context) error
	DeleteDifficulty(c echo.Context) error
}

type difficultyController struct {
	du usecase.IDifficultyUsecase
}

func NewDifficultyController(du usecase.IDifficultyUsecase) IDifficultyController {
	return &difficultyController{du}
}

// 公開エンドポイント（認証不要）

// 難易度を全て取得(非認証)
func (dc *difficultyController) GetAllDifficultiesPublic(c echo.Context) error {
	difficulties, err := dc.du.GetAllDifficulties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, difficulties)
}

// 認証が必要なエンドポイント

// 難易度を全て取得
func (dc *difficultyController) GetAllDifficulties(c echo.Context) error {
	difficulties, err := dc.du.GetAllDifficulties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, difficulties)
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

// 難易度を削除
func (dc *difficultyController) DeleteDifficulty(c echo.Context) error {
	id := c.Param("id")
	difficultyIdInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid difficulty id")
	}

	if difficultyIdInt < 0 {
		return c.JSON(http.StatusBadRequest, "difficulty id must be a positive integer")
	}

	difficultyIdUint := uint(difficultyIdInt)

	if err := dc.du.DeleteDifficulty(difficultyIdUint); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}