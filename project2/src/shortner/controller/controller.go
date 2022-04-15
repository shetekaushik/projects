package controller

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"os"
	obj "projects/project2/sqlconnect"
	"projects/project2/src/shortner/model"

	"github.com/itchyny/base58-go"
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
		shortenUrl = GenerateShortLink(longURL.URL)
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

func GenerateShortLink(longURL string) string {
	urlHashBytes := sha256Of(longURL)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
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
