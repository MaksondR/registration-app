package helper

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"registration-app/model"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"message": msg})
}

func RespondWithJWT(w http.ResponseWriter, code int, payload interface{}){
	jsonPayload, _ := json.Marshal(payload)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"code" : code,
	})

	tokenString, _ := token.SignedString([]byte("key"))

	json.NewEncoder(w).Encode(model.JwtToken{Token:tokenString})

	w.Write(jsonPayload)
}