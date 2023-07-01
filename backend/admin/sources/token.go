package sources

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"os"
)

func GenerateToken(user_id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("secret_key")))
}