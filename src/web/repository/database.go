package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	userName := "root"
	passWd := "root"
	url := "db:3306"
	dataBase := "sample"

	var err error
	dsn := fmt.Sprintf("%s:%s@(%s)/%s", userName, passWd, url, dataBase)
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
