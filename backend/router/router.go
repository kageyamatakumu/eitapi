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

	// 英単語
	wb := api.Group("/word-books")
	wb.GET("/", wbc.GetAllWordBooks)
	wb.POST("/", wbc.CreateWordBook)
	wb.GET("/:wordBookID/words", wc.GetAllWordsForWordBook)

	return e
}
