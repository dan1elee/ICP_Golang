package token

import (
	"ICP_Golang/enums"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Username   string `json:"username"`
	UserAccess uint   `json:"useraccess"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("ICP-secret-key")

func GenToken(username string, useraccess uint) (string, error) {
	c := MyClaims{
		username,
		useraccess,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "ICP-project",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

func HaveAccess(myClaims *MyClaims, level uint) bool {
	return myClaims.UserAccess == enums.STUDENT ||
		myClaims.UserAccess == enums.TEACHER ||
		myClaims.UserAccess == enums.ADMIN ||
		myClaims.UserAccess&level != 0
}
