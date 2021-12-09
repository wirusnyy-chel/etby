package middlewares

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/golang-jwt/jwt"
)

// var signMethod = jwt.GetSigningMethod("RS256")
var publicKey *rsa.PublicKey

type Claims struct {
	ExpiresAt int `json:"exp,omitempty"`
	Id        int `json:"jti,omitempty"`
}

func init() {
	fmt.Println("JWT initialization...")
	key, err := ioutil.ReadFile("demo.rsa.pub")
	if err != nil {
		panic(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		panic(err)
	}
}

func parseToken(tokenString string) (*jwt.Token, *jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	return token, &claims, err
}
