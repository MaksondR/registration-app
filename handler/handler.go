package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"registration-app/driver"
	"registration-app/helper"
	"registration-app/model"
	"registration-app/repository"
	"registration-app/repository/user"
)

func NewUserHandler(db *driver.DB) *User {
	return &User{
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
		helper.RespondWithJWT(w, user, map[string]string{"message": "Successfully Created"})
	}
}

func (u *User) UpdateRole(w http.ResponseWriter, r *http.Request) {

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
