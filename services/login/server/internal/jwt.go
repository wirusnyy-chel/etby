package internal

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

var signMethod = jwt.GetSigningMethod("RS256")
var secretKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var expPer = time.Hour / 2

func init() {

	var err error
	log.Info("JWT initialization...")
	key, err := ioutil.ReadFile("demo.rsa")
	if err != nil {
		log.Panic(err)
	}
	secretKey, err = jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Panic(err)
	}
	key, err = ioutil.ReadFile("demo.rsa.pub")
	if err != nil {
		log.Panic(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		log.Panic(err)
	}
}

func generateToken(userId int) (authToken, refreshToken string, err error) {
	token := jwt.NewWithClaims(signMethod, jwt.StandardClaims{
		Id:        fmt.Sprint(userId),
		ExpiresAt: time.Now().Add(expPer).Unix(),
	})
	authToken, err = token.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}
	token = jwt.NewWithClaims(signMethod, jwt.StandardClaims{
		Id:        fmt.Sprint(userId),
		ExpiresAt: time.Now().Add(expPer * 4).Unix(),
	})
	refreshToken, err = token.SignedString(secretKey)
	return
}
func parseToken(tokenString string) (int, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return 0, err
	}
	var exp float64
	var jti int
	for key, val := range claims {
		switch key {
		case "exp":
			exp, err = strconv.ParseFloat(fmt.Sprint(val), 64)
		case "jti":
			jti, err = strconv.Atoi(fmt.Sprint(val))
		}
	}
	if err != nil {
		return 0, err
	}
	if token.Valid {
		if time.Now().After(time.Unix(int64(exp), 0)) {
			return 0, fmt.Errorf("token expired")
		}
		return jti, nil
	}
	return 0, fmt.Errorf("token error")
}
