package controller

import (
	"errors"
	"fmt"
	obj "projects/project2/sqlconnect"
	linkshortner "projects/project2/src/Linkshortner"
	"projects/project2/src/shortner/model"
)

var db = obj.InitSQL()

func CreateShortUrl(longURL model.ShortLongURL) (error, *model.ShortLongURL) {

	row, err := db.Query("CALL `isExist`(?)", longURL.URL)
	if err != nil {
		return err, nil
	}
	defer row.Close()
	var isExist int
	if row.Next() {
		err = row.Scan(&isExist)
		if err != nil {
			return err, nil
		}
	}
	var shortURL model.ShortLongURL
	var shortenUrl string
	fmt.Println("isExist:", isExist)
	host := "http://localhost:9000/"
	if isExist == 0 {
		shortenUrl = linkshortner.GenerateShortLink(longURL.URL)
		err := SaveUrlMapping(shortenUrl, longURL.URL)
		if err != nil {
			return err, nil
		}
		shortURL.URL = host + shortenUrl
	} else {
		_ = row.NextResultSet()
		_ = row.Next()
		err = row.Scan(&shortenUrl)
		if err != nil {
			return err, nil
		}
		shortURL.URL = host + shortenUrl
	}
	return nil, &shortURL
}

func SaveUrlMapping(shortURL, LongURL string) error {
	row, err := db.Query("CALL `AddLongAndShortenURL`(?, ?)", shortURL, LongURL)
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

func ConShortUrlRedirect(shortURL string) (error, string) {
	row, err := db.Query("CALL `GetLongURL`(?)", shortURL)
	if err != nil {
		return err, ""
	}
	defer row.Close()
	var LongURL string
	if row.Next() {

		err = row.Scan(&LongURL)
		if err != nil {
			return err, ""
		}
		return nil, LongURL
	}
	errMsg := "No Record Found"
	return errors.New(errMsg), ""
}
