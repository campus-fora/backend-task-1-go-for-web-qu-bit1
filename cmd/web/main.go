package main

import (
	"campus_fora_week1/pkg/store"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func main() {

	// creating a database connection
	db := store.Conn()

	r := gin.Default()
	r.GET("/fetch", func(c *gin.Context) {
		finalData := store.GetAll(db)
		c.JSON(200, finalData)
	})

	r.POST("/add", func(c *gin.Context) {

		id := c.PostForm("id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id = %v", i)
		post := c.PostForm("post")
		fmt.Printf("id: %d; post: %s", i, post)
		_, err = db.Exec("INSERT INTO posts(ID,post) values(?,?)", i, post)
		if err != nil {
			log.Fatal(err)
		}

	})

	r.DELETE("/delete", func(c *gin.Context) {
		id := c.PostForm("id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec("DELETE from posts where id = ?", i)
		if err != nil {
			log.Fatal(err)
		}

	})
	r.Run()

}
