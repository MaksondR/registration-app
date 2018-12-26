package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func JwtValidate(token string) bool{
	tk, _ := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Incorrect signing method!")
		}
		return []byte("secret"), nil
	})
	if tk.Valid {
		return true
	} else {
		return false
	}

}

