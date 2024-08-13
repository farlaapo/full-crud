package middlewares

import (
	"full-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authanticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorization"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorization", "errors": err.Error()})
		return
	}
	context.Set("userId", userId)
	context.Next()

}
