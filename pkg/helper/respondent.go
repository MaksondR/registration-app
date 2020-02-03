package helper

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"registration-app/pkg/profile/model"
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

func RespondWithJWT(w http.ResponseWriter, user model.User, payload interface{}){
	jsonPayload, _ := json.Marshal(payload)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login" : user.Login,
		"password" : user.Login,
	})
	tokenString, _ := token.SignedString([]byte("key"))

	json.NewEncoder(w).Encode("{token: " + tokenString + "}")
	w.Write(jsonPayload)
}