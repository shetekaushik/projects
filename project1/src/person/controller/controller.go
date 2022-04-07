package controller

import (
	"errors"
	obj "projects/project1/sqlconnect"
	"projects/project1/src/person/model"
)

var db = obj.InitSQL()

func ConGetPersonByID(id string) (error, *model.PersonInfo) {
	row, err := db.Query("CALL `GetPersonByID`(?)", id)
	if err != nil {
		return err, nil
	}
	defer row.Close()
	if row.Next() {
		var PersonInfo model.PersonInfo
		err = row.Scan(&PersonInfo.Name, &PersonInfo.PhoneNumber, &PersonInfo.City,
			&PersonInfo.State, &PersonInfo.Street1, &PersonInfo.Street2,
			&PersonInfo.ZipCode)
		if err != nil {
			return err, nil
		}
		return nil, &PersonInfo
	}
	errMsg := "No Record Found"
	return errors.New(errMsg), nil
}
