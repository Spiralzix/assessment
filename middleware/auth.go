package middleware

import (
	"net/http"

	"github.com/Spiralzix/assessment/entity"
	"github.com/labstack/echo"
)

func Authorization(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header != "November 10, 2009" {
			return c.JSON(http.StatusUnauthorized, entity.Err{Message: "Unauthorized"})
		}
		return handler(c)
	}
}
