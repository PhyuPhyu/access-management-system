package middleware

import (
	"access-management-system/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware to access only by admin role
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")

		token, err := VerifyToken(bearerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			roleId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["role"]), 10, 64)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
				return
			}
			if roleId != 1 {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
				return
			}

		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
			return
		}
		ctx.Next()
	}
}

// Extract jwt token
func ExtractToken(bearerToken string) string {
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// Verfiy jwt token
func VerifyToken(bearerToken string) (*jwt.Token, error) {
	config, _ := config.LoadConfig(".")
	tokenString := ExtractToken(bearerToken)
	if tokenString == "" {
		return nil, errors.New("Authorization token required")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWTSecretToken), nil
	})

	if err != nil {
		if err.Error() == "Token is expired" {
			return nil, err
		} else {
			return nil, errors.New("Token Invalid")
		}
	}
	return token, nil
}
