package main

import (
	"github.com/gin-gonic/gin"
)

type PostData struct {
	UID  string `json:"uid"`
	Area string `json:"area"`
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H {"message": "pong",} )
}

func postDataHandler(c *gin.Context) {
	var postData PostData

	if err := c.BindJSON(&postData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	responseData := gin.H{
		"uid":  postData.UID,
		"area": postData.Area,
	}

	c.JSON(200, responseData)
}

func handleRequests() {
	r := gin.Default()
	r.GET("/ping", PingHandler)
	
	r.POST("/post/m5", postDataHandler)

	r.Run()
}

func main() {
	handleRequests()
}
