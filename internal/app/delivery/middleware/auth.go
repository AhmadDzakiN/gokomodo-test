package middleware

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/pkg/jwt"
	"net/http"
	"os"
	"strings"
)

func JWTAuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Not in the login session, please login again"})
			return
		}

		data, err := jwt.ValidateToken(token, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error": err.Error()})
			return
		}

		tokenUserData := data.(jwt.JWTCustomClaims)
		ctx.Set("token", tokenUserData)
		ctx.Next()
	}
}

func RoleAuthorization(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("token")
		if claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found"})
			c.Abort()
			return
		}

		customClaims, ok := claims.(jwt.JWTCustomClaims)
		if !ok || customClaims.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
