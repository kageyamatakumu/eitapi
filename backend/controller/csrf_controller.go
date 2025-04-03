package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ICsrfTokenController interface {
	CsrfToken(c echo.Context) error
}

type csrfTokenController struct{}

func NewCsrfTokenController() ICsrfTokenController {
	return &csrfTokenController{}
}

func (cc *csrfTokenController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
