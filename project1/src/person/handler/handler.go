package handler

import (
	"net/http"
	obj "projects/project1/sqlconnect"

	con "projects/project1/src/person/controller"

	"github.com/labstack/echo"
)

var db = obj.InitSQL

func APIGetPersonInfo(echo echo.Context) error {

	id := echo.Param("id")

	err, personInfo := con.ConGetPersonByID(id)
	if err != nil {
		return err
	}
	return echo.JSON(http.StatusOK, personInfo)
}
