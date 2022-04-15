package sqlconnect

import (
	"database/sql"
	"projects/project2/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db1 *sql.DB
var once sync.Once

func InitSQL() *sql.DB {
	once.Do(func() {
		d, err := sql.Open("mysql", config.USER+":"+config.PASSWORD+"@tcp("+config.HOST+
			":"+config.PORT+")/"+config.DB+"?charset=utf8&parseTime=True")
		if err != nil {
			panic(err)
		} else {
			db1 = d
			d.SetConnMaxIdleTime(10)
			d.SetMaxOpenConns(50)
		}
	})
	return db1
}
