package router

import (
	"backend/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(ac controller.IAdminController, wbc controller.IWordBookController, wc controller.IWordController) *echo.Echo {
	e := echo.New()

	v1 := e.Group("/api/v1")

	api := v1.Group("")
	api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	// 管理者
	a := v1.Group("/admins")
	a.POST("/", ac.CreateAdmin)
	a.POST("/login", ac.LoginAdmin)

	// 英単語帳
	wb := api.Group("/word-books")
	wb.GET("/", wbc.GetAllWordBooks)
	wb.POST("/", wbc.CreateWordBook)
	wb.DELETE("/:wordBookId", wbc.DeleteWordBook)
	// 英単語
	wb.GET("/:wordBookId/words", wc.GetAllWordsForWordBook)
	wb.GET("/:wordBookId/words/:wordId", wc.GetWordById)
	wb.POST("/:wordBookId/words", wc.CreateMultipleWords)
	wb.PUT("/:wordBookId/words/:wordId", wc.UpdateWord)
	wb.DELETE("/:wordBookId/words/:wordId", wc.DeleteWord)

	return e
}
