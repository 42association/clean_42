package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type PostData struct {
	UID  string `json:"uid"`
	Place string `json:"place"`
}

func insertCleanedData(db *sql.DB, data PostData) {
	_, err := db.Exec("INSERT INTO table_clean (uid, place) VALUES (?, ?)", data.UID, data.Place,)
	if (err != nil) {
		log.Println(err)
	}
}

func openMariadb() *sql.DB {
	db, err := sql.Open("mysql", "db-user:db-pass@tcp(db:3306)/db_clean")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()
	return db
}

func postDataHandler(c *gin.Context) {
	var postData PostData

	if err := c.BindJSON(&postData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		log.Println(err)
		return
	}
	db := openMariadb()
	if db == nil {
		return
	}
	log.Println(postData.UID, postData.Place)
	insertCleanedData(db, postData)
}

func handleRequests() {
	r := gin.Default()
	
	r.POST("/post/m5", postDataHandler)

	r.Run()
}

func main() {
	handleRequests()
}
