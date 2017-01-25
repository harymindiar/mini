package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

type Users struct{}

func (u *Users) Add(e *echo.Echo) *echo.Echo {
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is users resource")
	})

	return e
}
