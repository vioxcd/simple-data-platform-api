package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	RoleId uint   `json:"roleId"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, name string, roleId uint) string {
	claims := &JwtCustomClaims{
		userId,
		name,
		roleId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return t
}
