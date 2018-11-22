package auth

import (
	"crypto"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

type JwtAdapter struct {
	privateKey crypto.PrivateKey
	publicKey crypto.PublicKey
}

func NewJwtAdapter() *JwtAdapter {
	pemKeyBytes, err := ioutil.ReadFile("../data/key.pem")
	if err != nil {panic(err)}

	pem, err := jwt.ParseRSAPrivateKeyFromPEM(pemKeyBytes)
	if err != nil {panic(err)}

	pubKeyBytes, err := ioutil.ReadFile("../data/key.pub")
	if err != nil {panic(err)}

	pub, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {panic(err)}

	return &JwtAdapter{
		privateKey: pem,
		publicKey: pub,
	}
}

func (j *JwtAdapter) Generate(i *Identity) (string, error) {

	claims := jwt.MapClaims{
		"identity": i,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

func (j *JwtAdapter) Validate(signedStr string) error {

	t, err := j.fetchToken(signedStr)
	if err != nil {
		return err
	}

	if !t.Valid {
		return errors.New("token is not valid")
	}
	return nil
}

func (j *JwtAdapter) fetchToken(signedStr string) (*jwt.Token, error) {
	return jwt.Parse(signedStr, func(t *jwt.Token) (interface{}, error){
		return j.publicKey, nil
	})
}

func (j *JwtAdapter) Read(signedStr string) (*Identity, error) {
	t, err := j.fetchToken(signedStr)
	if err != nil {
		return nil, err
	}

	mClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}

	i, ok := mClaims["identity"]
	if !ok {
		return nil, errors.New("invalid claims map")
	}

	im, ok := i.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid identity type")
	}

	id, ok := im["Id"].(float64)
	if !ok {
		return nil, errors.New("invalid id type")
	}

	name, ok := im["Name"].(string)
	if !ok {
		return nil, errors.New("invalid name type")
	}

	return &Identity{Id: int(id), Name: name}, nil
}
