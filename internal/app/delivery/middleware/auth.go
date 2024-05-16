package middleware

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/pkg/jwt"
	"net/http"
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "status_code": http.StatusUnauthorized, "error": "Not in the login session, please login again"})
			ctx.Abort()
			return
		}

		data, err := jwt.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "status_code": http.StatusUnauthorized, "error": "Failed to validate token"})
			ctx.Abort()
			return
		}

		tokenUserData := data.(jwt.JWTCustomClaims)
		ctx.Set("token", tokenUserData)
		ctx.Next()
	}
}

func RoleAuthorization(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("token")
		if claims == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "error", "status_code": http.StatusUnauthorized, "error": "No token claims found"})
			ctx.Abort()
			return
		}

		customClaims, ok := claims.(jwt.JWTCustomClaims)
		if !ok || customClaims.Role != role {
			ctx.JSON(http.StatusForbidden, gin.H{"status": "error", "status_code": http.StatusUnauthorized, "error": "Access forbidden"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
