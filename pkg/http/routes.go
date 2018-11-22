package http

import (
	"github.com/gin-gonic/gin"
)



func signin(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": "t",
	})
}

func signout(c *gin.Context) {

}

func signup(c *gin.Context) {

}
