package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"ril.api-ia/internal/application/usecase"
)

func AuthMiddleware(userUseCase usecase.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiAiToken := GetBearerToken(c)
		if apiAiToken == "" {
			log.Print("apiAiToken is empty")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		user, err := userUseCase.GetUserByApiAiToken(apiAiToken)
		if err != nil {
			log.Print(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		if user == nil {
			log.Print("user is nil")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		c.Set("user", user)
		c.Next()
		//TODO: Lanzar evento de autenticacion
	}
}

func GetBearerToken(c *gin.Context) string {
	auth := strings.TrimSpace(c.GetHeader("Authorization"))
	if auth == "" {
		return ""
	}

	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}

	return parts[1]
}
