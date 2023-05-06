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
	//db.Query returns pointer to rows
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	//making a variable data of type posts can also be made inside the loop
	// sole purpose of data variable is as a temporary variable to append into the final array
	var data Posts
	// storing final data into an array and iterating through it
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
