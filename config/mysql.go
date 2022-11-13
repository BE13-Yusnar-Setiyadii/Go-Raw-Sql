package config

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectToDB() *sql.DB {
	var connectionString = "root:yusnar12345@tcp(127.0.0.1:3306)/yusnarsetiyadi"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connection to db", err.Error())
	} else {
		fmt.Println("connection success")
	}

	return db
}
