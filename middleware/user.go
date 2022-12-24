package middleware

import (
	"access-management-system/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")

		token, err := VerifyToken(bearerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		var user models.User
		if ok && token.Valid {
			userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userid"]), 10, 64)
			if err != nil {
				fmt.Println("err", err.Error())
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

			err = models.GetUser(&user, int(userId))
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
				return
			}

			if !user.IsAllowByAdmin {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Your account need to be approved by admin"})
				return
			}

		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
