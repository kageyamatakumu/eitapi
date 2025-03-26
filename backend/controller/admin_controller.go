package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IAdminController interface {
	CreateAdmin(c echo.Context) error
	LoginAdmin(c echo.Context) error
}

type adminController struct {
	au usecase.IAdminUsecase
}

func NewAdminController(au usecase.IAdminUsecase) IAdminController {
	return &adminController{au}
}

// 管理者を新規作成
func (ac *adminController) CreateAdmin(c echo.Context) error {
	admin := model.Admin{}
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newAdmin, err := ac.au.CreateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newAdmin)
}

// 管理者のログイン
func (ac *adminController) LoginAdmin(c echo.Context) error {
	admin := model.Admin{}
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := ac.au.LoginAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}