package router

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(wbc controller.IWordBookController, wc controller.IWordController) *echo.Echo {
	e := echo.New()

	// 英単語
	wb := e.Group("/word_books")
	wb.GET("/", wbc.GetAllWordBooks)
	wb.GET("/:wordBookID/words", wc.GetAllWordsForWordBook)

	return e
}
