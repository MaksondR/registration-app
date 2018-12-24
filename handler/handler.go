package handler

import (
	"encoding/json"
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
		helper.RespondWithJWT(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

