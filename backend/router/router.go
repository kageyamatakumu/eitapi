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
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
	}))

	v1 := e.Group("/api/v1")

	v1.GET("/csrf", cc.CsrfToken)

	// 非認証エンドポイント
	unauthorized := v1.Group("/unauthorized")

	// 管理者用(非認証)
	unauthorizedAdmin := unauthorized.Group("/admins")
	unauthorizedAdmin.POST("/login", ac.LoginAdmin)

	// 英単語帳(非認証)
	unauthorizedWordBook := unauthorized.Group("/word-books")
	unauthorizedWordBook.GET("/", wbc.GetAllWordBooksPublic)

	// 英単語(非認証)
	unauthorizedWord := unauthorizedWordBook.Group("/:wordBookId/words")
	unauthorizedWord.GET("/:id", wc.GetAllWordsForWordBookPublic)

	// 認証が必要なエンドポイント
	authorizedApi := v1.Group("/authorized")
	authorizedApi.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	// 管理者用エンドポイント
	authorizedAdmin := authorizedApi.Group("/admins")
	authorizedAdmin.POST("/", ac.CreateAdmin)

	// 管理者用 英単語帳エンドポイント
	authorizedAdminWordBook := authorizedAdmin.Group("/word-books")
	authorizedAdminWordBook.GET("/", wbc.GetAllWordBooks)
	authorizedAdminWordBook.POST("/", wbc.CreateWordBook)
	authorizedAdminWordBook.DELETE("/:wordBookId", wbc.DeleteWordBook)

	// 管理者用 英単語エンドポイント
	authorizedAdminWord := authorizedAdminWordBook.Group("/:wordBookId/words")
	authorizedAdminWord.GET("/", wc.GetAllWordsForWordBook)
	authorizedAdminWord.GET("/:id", wc.GetWordById)
	authorizedAdminWord.POST("/", wc.CreateMultipleWords)
	authorizedAdminWord.PUT("/:id", wc.UpdateWord)
	authorizedAdminWord.DELETE("/:id", wc.DeleteWord)

	return e
}
