package handler

import (
	"fmt"
	"net/http"
	obj "projects/project2/sqlconnect"

	con "projects/project2/src/shortner/controller"
	"projects/project2/src/shortner/model"

	"github.com/labstack/echo"
)

var db = obj.InitSQL

func APIGetShortenLink(echo echo.Context) error {
	var url model.ShortLongURL
	if err := echo.Bind(&url); err != nil {
		return err
	}

	err, shortenURL := con.CreateShortUrl(url)
	if err != nil {
		return err
	}
	return echo.JSON(http.StatusOK, shortenURL)
}

func APIShortUrlRedirect(echo echo.Context) error {
	shortUrl := echo.Param("shortUrl")
	err, longURL := con.ConShortUrlRedirect(shortUrl)
	if err != nil {
		return err
	}
	fmt.Println("longURL:", longURL)
	return echo.Redirect(302, longURL)
}
