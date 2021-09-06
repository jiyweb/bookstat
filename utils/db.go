package utils

import (
	"database/sql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err := sql.Open("mysql", "bool:123456@tcp(127.0.0.1:3306)/bool")
	if err != nil {

		panic(err.Error())
	}
}
