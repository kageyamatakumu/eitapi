package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IWordController interface {
	GetAllWordsForWordBook(c echo.Context) error
	CreateWord(c echo.Context) error
}

type wordController struct {
	wu usecase.IWordCase
}

func NewWordController(wu usecase.IWordCase) IWordController {
	return &wordController{wu}
}

// 英単語帳に紐づく英単語を全て取得
func (wc *wordController) GetAllWordsForWordBook(c echo.Context) error {
	id := c.Param("wordBookID")
	wordBookIDInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid wordBookID")
	}
	wordBookIDUint := uint(wordBookIDInt)

	words, err := wc.wu.GetAllWordsForWordBook(wordBookIDUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, words)
}

// 英単語を新規作成
func (wc *wordController) CreateWord(c echo.Context) error {
	word := model.Word{}
	if err := c.Bind(&word); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	wordRes, err := wc.wu.CreateWord(word);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, wordRes)
}
