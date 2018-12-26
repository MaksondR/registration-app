package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"registration-app/driver"
	"registration-app/model"
	"registration-app/repository"
	"registration-app/repository/user"
	"registration-app/helper"
)

func NewUserHandler(db *driver.DB) *User{
	return &User {
		repo: user.NewSQLUserRepo(db.SQL),
	}
}

type User struct {
	repo repository.UserRepo
}

func (u *User) RegisterUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{}

	error := json.Unmarshal(reqBody, &user)
	if error != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, "Unmarshal JSON Error!")
	}

	err = u.repo.RegisterUser(r.Context(), &user)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		helper.RespondWithJWT(w, map[string]string{"message": "Successfully Created"})
	}
}

func (u *User) UpdateRole(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	if helper.JwtValidate(token) {
		helper.RespondWithError(w, http.StatusForbidden, "Invalid authorization token!")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{}

	error := json.Unmarshal(reqBody, &user)
	if error != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, "Unmarshal JSON Error!")
	}

	err = u.repo.UpdateAdmin(r.Context(), &user)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Role successfully changed!"})
	}
}

