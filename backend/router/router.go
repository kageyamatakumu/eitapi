package router

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(wc controller.IWordBookController) *echo.Echo {
	e := echo.New()

	// 英単語帳
	wb := e.Group("/word_books")
	wb.GET("/", wc.GetAllWordBooks)

	return e
}
