package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IWordBookController interface {
	GetAllWordBooks(c echo.Context) error
	CreateWordBook(c echo.Context) error
	DeleteWordBook(c echo.Context) error
}

type wordBookController struct {
	wu usecase.ITWordBookCase
}

func NewWordBookController(wu usecase.ITWordBookCase) IWordBookController {
	return &wordBookController{wu}
}

// 英単語帳を全て取得
func (wc *wordBookController) GetAllWordBooks(c echo.Context) error {
	wordBooks, err := wc.wu.GetAllWordBook()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, wordBooks)
}

// 英単語帳を新規作成
func (wc *wordBookController) CreateWordBook(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	wordBook := model.WordBook{}
	if err := c.Bind(&wordBook); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	wordBook.AdminId = uint(userId.(float64))
	wordBookRes, err := wc.wu.CreateWordBook(wordBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, wordBookRes)
}

// 英単語帳を削除
func (wc *wordBookController) DeleteWordBook(c echo.Context) error {
	id := c.Param("wordBookId")
	wordBookId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "word book id must be an integer")
	}

	if wordBookId < 0 {
		return c.JSON(http.StatusBadRequest, "word book id must be a positive integer")
	}

	wordBookIdUint := uint(wordBookId)

	if err := wc.wu.DeleteWordBook(wordBookIdUint); err != nil {
		if err.Error() == "object does not exist" {
			return c.JSON(http.StatusNotFound, "word book not found")
		}
		return c.JSON(http.StatusInternalServerError, "failed to delete word book")
	}

	return c.NoContent(http.StatusNoContent)
}
