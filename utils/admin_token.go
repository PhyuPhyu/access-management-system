package utils

import (
	"access-management-system/config"
	"access-management-system/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/outbrain/golib/log"
)

// Create admin token
func CreateAdminToken(admin models.Admin) (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["admin_name"] = admin.Name
	atClaims["role"] = admin.RoleId
	atClaims["exp"] = time.Now().AddDate(0, 6, 0).Unix()

	config, _ := config.LoadConfig(".")

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(config.JWTSecretToken))
	if err != nil {
		log.Error("Sign token error: ", err)
		return "", err
	}
	return token, nil
}
