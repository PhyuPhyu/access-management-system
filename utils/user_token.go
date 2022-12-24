package utils

import (
	"access-management-system/config"
	"access-management-system/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/outbrain/golib/log"
)

// Create JWT token
func CreateToken(user models.User) (token string, err error) {
	// Token expire time is 2 hours
	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = user.ID
	claims["name"] = user.Name
	claims["exp"] = expirationTime.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	config, _ := config.LoadConfig(".")

	token, err = at.SignedString([]byte(config.JWTSecretToken))
	if err != nil {
		log.Error("Sign token error: ", err)
		return "", err
	}

	return token, nil
}
