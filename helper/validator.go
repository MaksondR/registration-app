package helper

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func JwtValidate(token string) bool {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Incorrect signing method!")
		}
		return []byte("key"), nil
	})
	if err != nil {
		return false
	}
	if tk.Valid {
		return true
	} else {
		return false
	}

}
