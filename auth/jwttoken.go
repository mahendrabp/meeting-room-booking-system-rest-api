package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

//CreateToken: create jwt token
func CreateToken(id uint, role string) (string, error) {
	claims := jwt.MapClaims{}
	fmt.Println(id, role)
	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(os.Getenv("API_SECRET"), "ssst this is a secret")
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
