package main

import (
	"campus_fora_week1/pkg/store"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func main() {

	// creating a database connection
	// Conn function made in store.go
	db := store.Conn()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	// making a get request
	r.GET("/fetch", func(c *gin.Context) {
		// getAll function made in store.go
		finalData := store.GetAll(db)
		c.JSON(200, finalData)
	})

	// making a post request
	r.POST("/add", func(c *gin.Context) {

		id := c.PostForm("id")
		// post form returns a string value
		// so we convert string to int using parseint
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

	// making a delete request
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
	err := r.Run()
	if err != nil {
		return
	}

}
