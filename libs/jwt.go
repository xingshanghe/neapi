package libs

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 自定义声明
type NeJwtClaims struct {
	Account interface{} `json:"account"`
	jwt.StandardClaims
}

func CreateJwt(account interface{}) (string, error) {
	claims := NeJwtClaims{
		account,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("ne----admin"))
}
