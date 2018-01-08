package libs

import (
	"time"

	"github.com/astaxie/beego/config"
	jwt "github.com/dgrijalva/jwt-go"
)

// 自定义声明
type NeJwtClaims struct {
	Account interface{} `json:"account"`
	Roles   []string    `json:"roles"`
	jwt.StandardClaims
}

func CreateJwt(account interface{}, roles []string) (string, error) {
	claims := NeJwtClaims{
		account,
		roles,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	conf, _ := config.NewConfig("ini", "conf/app.conf")

	return token.SignedString([]byte(conf.String("salt")))
}
