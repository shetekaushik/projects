package route

import (
	handler "projects/project2/src/shortner/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitServer() *echo.Echo {
	e := echo.New()
	url := e.Group("/api/shorten/link")
	url1 := e.Group("")

	e.Use(middleware.Logger())

	url.POST("", handler.APIGetShortenLink)
	url1.GET("/:shortUrl", handler.APIShortUrlRedirect)

	return e
}
