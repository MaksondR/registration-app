package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	authHandler "registration-app/pkg/auth/handler"
	"registration-app/pkg/profile/driver"
	"registration-app/pkg/profile/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connection, err := driver.ConnectSQL("root:root@tcp(127.0.0.1:3306)/db?charset=utf8")
	if err != nil {
		panic(err)
	}

	go runProfileApp(connection)
	runAuthApp()
}

func runProfileApp(dbConn *driver.DB) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	uHandler := handler.NewUserHandler(dbConn)

	router.Post("/users", uHandler.RegisterUser)
	router.Put("/api/admins", uHandler.UpdateRole)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func runAuthApp() {
	err := rpc.Register(&authHandler.Handler{})

	if err != nil {
		log.Fatal("Register error: ", err)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
