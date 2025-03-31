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
	GetWordById(c echo.Context) error
	CreateWord(c echo.Context) error
	UpdateWord(c echo.Context) error
	DeleteWord(c echo.Context) error
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

// 英単語を取得
func (wc *wordController) GetWordById(c echo.Context) error {
	id := c.Param("wordID")
	wordIdInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid wordID")
	}
	wordIdUint := uint(wordIdInt)

	word, err := wc.wu.GetWordById(wordIdUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, word)
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

// 英単語を更新
func (wc *wordController) UpdateWord(c echo.Context) error {
	id := c.Param("wordID")
	wordIDInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid wordBookID")
	}
	wordIDUint := uint(wordIDInt)

	word := model.Word{}
	if err := c.Bind(&word); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	wordRes, err := wc.wu.UpdateWord(word, wordIDUint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, wordRes)
}

// 英単語を削除
func (wc *wordController) DeleteWord(c echo.Context) error {
	id := c.Param("wordId");
	wordIdInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wordId must be an integer")
	}

	if wordIdInt < 0 {
		return c.JSON(http.StatusBadRequest, "wordId must be a positive integer")
	}

	wordUint := uint(wordIdInt)

	if err := wc.wu.DeleteWord(wordUint); err != nil {
		if err.Error() == "object does not exist" {
			return c.JSON(http.StatusNotFound, "word not found")
		}
		return c.JSON(http.StatusInternalServerError, "failed to delete word")
	}

	return c.NoContent(http.StatusNoContent)
}
