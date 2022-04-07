package route

import (
	handler "projects/project1/src/person/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitServer() *echo.Echo {
	e := echo.New()
	person := e.Group("/api/person/")

	e.Use(middleware.Logger())

	person.GET(":id/get", handler.APIGetPersonInfo)

	return e
}
