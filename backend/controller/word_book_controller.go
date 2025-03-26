package controller

import (
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IWordBookController interface {
	GetAllWordBooks(c echo.Context) error
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
