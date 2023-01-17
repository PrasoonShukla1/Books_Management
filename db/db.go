package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connecttomysql() *sql.DB {
	db, err := sql.Open("mysql", "root:Prasoon$123@tcp(localhost:3306)/library")
	if err != nil {
		panic(err)
	}
	return db
}
