package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "106.52.121.11"
	port     = 8081
	user     = "postgres"
	password = "postgres"
	dbname   = "ard"
)

func DBinit() {
	var err error
	//打开数据库
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)

	err = db.Ping()
	checkErr(err)
	fmt.Println("Successfully connected!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Close() {
	db.Close()
}
