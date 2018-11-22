package main

import (
	"auth/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
