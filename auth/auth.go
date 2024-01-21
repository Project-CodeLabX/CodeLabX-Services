package auth

import (
	"codelabx/models"
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
