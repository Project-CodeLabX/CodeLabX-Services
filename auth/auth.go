package auth

import (
	"codelabx/models"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secret_key = []byte("CodeLabX_Rocks")
)

type claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) string {
	claims := &claims{
		UserName: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(24 * time.Hour)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret_key)

	if err != nil {
		log.Fatal("error during create token : ", err)
		return "Not able to create token"
	}

	return tokenStr
}

func IsAuthorized(tokenStr string) bool {

	claims := &claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return secret_key, nil })
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("signature invalid...")
			return false
		}
		fmt.Println("Bad request...")
		return false
	}

	if !token.Valid {
		fmt.Println("Unauthorised token...")
		return false
	}
	fmt.Println("Authorized : valid token...")
	return true
}
