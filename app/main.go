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

func insertCleanedData(db *sql.DB, data PostData)int64 {
	res, err := db.Exec("INSERT INFO table_clean (uid, place) VALUES (?, ?)", data.UID, data.Place,)
	if (err != nil) {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("insertUser res.LastInsertId error err:%v", err)
	}
	return id
}

func postDataHandler(c *gin.Context) {
	var postData PostData

	log.Println("post data to API starting")
	if err := c.BindJSON(&postData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	log.Println("send data to DB starting")
	db, err := sql.Open("mysql", "db-user:db-pass@tcp(127.0.0.1:3306)/db_clean")
	if err != nil {
		panic(err.Error())
		return
	}
	defer db.Close()
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
