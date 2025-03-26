package router

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(ac controller.IAdminController, wbc controller.IWordBookController, wc controller.IWordController) *echo.Echo {
	e := echo.New()

	// 管理者
	a := e.Group("/admins")
	a.POST("/", ac.CreateAdmin)
	a.POST("/login", ac.LoginAdmin)

	// 英単語
	wb := e.Group("/word_books")
	wb.GET("/", wbc.GetAllWordBooks)
	wb.GET("/:wordBookID/words", wc.GetAllWordsForWordBook)

	return e
}
