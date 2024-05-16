package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type JWTCustomClaims struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

func CreateToken(userID, name, role string) (token string, err error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"name":    name,
		"exp":     now.Add(time.Minute * 43200).Unix(),
		"iat":     now.Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Err(err).Msg("Error creating token")
		return
	}

	return
}

func ValidateToken(token string) (data interface{}, err error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	extractedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error in token parsing")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		err = fmt.Errorf("Invalidate token: %w", err)
		return
	}

	if extractedToken == nil {
		err = errors.New("Invalid token")
		return
	}

	claims, ok := extractedToken.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("Token error")
		return
	}

	data = JWTCustomClaims{
		UserID: claims["user_id"].(string),
		Name:   claims["name"].(string),
		Role:   claims["role"].(string),
	}

	return
}

func GetTokenClaims(ctx *gin.Context) (claims JWTCustomClaims, err error) {
	token, _ := ctx.Get("token")
	claims, ok := token.(JWTCustomClaims)
	if !ok || token == nil {
		err = errors.New("Invalid token")
		return
	}

	return
}
