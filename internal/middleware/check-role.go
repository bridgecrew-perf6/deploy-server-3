package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xavimg/Turing/apituringserver/internal/helper"
	"github.com/xavimg/Turing/apituringserver/internal/service"
)

func CheckRole(checkRole service.UserService) gin.HandlerFunc {
	return func(context *gin.Context) {

		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, _ := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte("turingoffworld"), nil
		})

		claims := token.Claims.(jwt.MapClaims)
		id := claims["UserID"]

		typeUser := checkRole.CheckRole(id)

		if typeUser != "admin" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "not allowed")
		}

		context.Next()
	}

}
