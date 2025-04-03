package router

import (
	"backend/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(cc controller.ICsrfTokenController, ac controller.IAdminController, wbc controller.IWordBookController, wc controller.IWordController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
	}))

	v1 := e.Group("/api/v1")

	v1.GET("/csrf", cc.CsrfToken)

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
