package handler

import (
	"net/http"
	obj "projects/project1/sqlconnect"

	con "projects/project1/src/person/controller"
	"projects/project1/src/person/model"

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

func APIAddPersonInfo(echo echo.Context) error {
	var personInfo model.PersonInfo
	if err := echo.Bind(&personInfo); err != nil {
		return err
	}
	err := con.ConAddPerson(personInfo)
	if err != nil {
		return err
	}
	return echo.NoContent(http.StatusOK)
}
