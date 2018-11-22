package http

import "github.com/gin-gonic/gin"

func RunHttpServer() {
	r := gin.Default()

	r.Use(addHeaders)
	InitRoutes(r)
	r.Run(":8033")
}

func addHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func InitRoutes(engine *gin.Engine) {
	engine.POST("/sign-in", signin)
	engine.POST("/sign-out", signout)
	engine.POST("/sign-up", signup)
}

