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

func ConAddPerson(personInfo model.PersonInfo) error {

	row, err := db.Query("CALL `AddPersonInfo`(?, ?, ?, ?, ?, ?, ?)",
		personInfo.Name, personInfo.PhoneNumber, personInfo.City,
		personInfo.State, personInfo.Street1, personInfo.Street2,
		personInfo.ZipCode)
	if err != nil {
		return err
	}
	defer row.Close()
	var isSuccesful int
	if row.Next() {
		err = row.Scan(&isSuccesful)
		if err != nil {
			return err
		}
	}
	if isSuccesful == 1 {
		return nil
	} else {
		_ = row.NextResultSet() && row.Next()
		var errMsg string
		err = row.Scan(&errMsg)
		if err != nil {
			return err
		}
		return errors.New(errMsg)
	}
}
