package main

import (
	"github.com/gin-gonic/gin"
	"auth/http"
)

func main() {
	r := gin.Default()

	r.Use(addHeaders)
	http.InitRoutes(r)
	r.Run(":8033")
}

func addHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
