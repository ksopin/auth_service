package main

import (
	"auth/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	RunHttpServer()
}

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
	engine.POST("/authenticated", isAuthenticated)
}



func signup(c *gin.Context) {
	ok, err := app.GetAppService().SignUp()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": ok,
	})
}


func signin(c *gin.Context) {
	token, err := app.GetAppService().SignIn()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func signout(c *gin.Context) {
	err := app.GetAppService().SignOut()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{})
}


func isAuthenticated(c *gin.Context) {
	ok, err := app.GetAppService().IsAuthenticated()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": ok,
	})
}


