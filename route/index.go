package route

import (
	"github.com/labstack/echo"
	"net/http"
)

type Index struct {
	Base
}

func (i *Index) Add(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!!")
	})

	return e
}
