package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"registration-app/driver"
	"registration-app/handler"
	"registration-app/helper"
	"testing"

	"github.com/go-chi/chi"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO User").WithArgs(helper.GenerateId, "Maks", "12345678", "stalk@int.ua", "user").WillReturnError(nil)
	mock.ExpectCommit()

	payload := []byte(`{"login":"Maks", "password":"12345678", "email":"stalk@int.ua"}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	router := chi.NewRouter()

	dbConn := &driver.DB{}
	dbConn.SQL = db
	uHandler := handler.NewUserHandler(dbConn)

	router.HandleFunc("/users", uHandler.RegisterUser)
	router.ServeHTTP(rr, req)

	checkResponseCode(t, http.StatusOK, rr.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
