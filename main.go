package main

import (
	"log"
	"net/http"
	"registration-app/driver"
	"registration-app/handler"
	"registration-app/mdlwr"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connection, err := driver.ConnectSQL("root:root@tcp(127.0.0.1:3306)/db?charset=utf8")
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(mdlwr.AuthMiddleware)

	uHandler := handler.NewUserHandler(connection)

	router.Post("/users", uHandler.RegisterUser)
	router.Put("/api/admins", uHandler.UpdateRole)
	log.Fatal(http.ListenAndServe(":8080", router))
}
