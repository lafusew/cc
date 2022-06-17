package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lafusew/cc/controllers"
	"github.com/lafusew/cc/data"
	"github.com/lafusew/cc/middleware"
)

var port = 8080

func main() {
	db := data.Connect(true)

	uController := controllers.UserController {Db: db}
	authController := controllers.AuthController {Db: db}

	http.HandleFunc("/users/", middleware.Authentication(uController.HandleUsers))
	http.HandleFunc("/signup/", authController.HandleAuth)
	http.HandleFunc("/signin/", authController.HandleAuth)

	log.Printf("Server starting on port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
