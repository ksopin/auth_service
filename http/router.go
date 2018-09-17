package http

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"time"
	"os"
	"bytes"
		"crypto/rsa"
		)

func InitRoutes(engine *gin.Engine) {
	engine.POST("/signin", signin)
	engine.POST("/signout", signout)
	engine.POST("/signup", signup)

	engine.POST("/debug", debug)
}

func signin(c *gin.Context) {
	t, err := generateSignedJWT()

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": t,
	})
}

func debug(c *gin.Context) {

	r := &struct{
		Token string `json:"token"`
	}{}

	if err := c.BindJSON(r); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := jwt.Parse(r.Token, hToken)
	c.JSON(200, gin.H{
		"token": t,
		"err": err.Error(),
	})
}

func hToken(t *jwt.Token) (interface{}, error) {
	return t.Signature, nil
}

func signout(c *gin.Context) {

}

func signup(c *gin.Context) {

}



func ok(c *gin.Context) {
	c.Status(http.StatusOK)
}

func generateSignedJWT() (string, error) {
	token := generateJWT()
	pk, err := getPem()
	if err != nil {
		return "", err
	}
	return token.SignedString(pk)
}

func generateJWT() *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer: "auth.io",
		Id: "1",
	})
}

func getPem() (*rsa.PrivateKey, error) {
	file, err := os.Open("./data/key.pem")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	return jwt.ParseRSAPrivateKeyFromPEM(buf.Bytes())
}