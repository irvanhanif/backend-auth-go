package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"trigonal/backend-auth/api/database"
	"trigonal/backend-auth/api/helper"
	"trigonal/backend-auth/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func ValidateToken(ctx *gin.Context) {
	fmt.Println("get token from header")
	token := ctx.GetHeader("Authorization")
	if token == "" {
		helper.ERROR(ctx, http.StatusUnauthorized, "User belum melakukan login")
		return
	}

	tokenString := strings.Split(token, " ")[1]
	fmt.Println("parsing jwt")
	tokenValidate, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY_API")), nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("validate token")
	if claims, ok := tokenValidate.Claims.(jwt.MapClaims); ok {
		userId := (claims["user"].(map[string]interface{}))["ID"]
		id := userId.(string)

		fmt.Println("get user from id in token")
		var userToken models.User
		database.DB.Omit("password").First(&userToken, map[string]string{"id": id})
		if userToken.ID == uuid.Nil {
			helper.ERROR(ctx, http.StatusUnauthorized, "User tidak memiliki akses")
			return
		}

		fmt.Println("set user to request")
		ctx.Set("user", userToken)
		ctx.Next()
	} else {
		helper.ERROR(ctx, http.StatusUnauthorized, "User belum melakukan login")
		return		
	}
}