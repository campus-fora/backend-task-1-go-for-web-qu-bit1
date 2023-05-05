package store

import (
	"database/sql"
	"fmt"
	"log"
)

type Posts struct {
	Id   int64
	Post string
}

func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:campus_fora@tcp(localhost:3306)/campus_fora")
	if err != nil {
		fmt.Println("error validating sql.open")
		panic(err.Error())
	}
	return db
}

func GetAll(db *sql.DB) []Posts {
	res, err := db.Query("SELECT * FROM posts")
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()
	var data Posts
	var finalData []Posts
	for res.Next() {
		err := res.Scan(&data.Id, &data.Post)
		if err != nil {
			log.Fatal(err)
		}
		finalData = append(finalData, data)
	}
	return finalData
}
