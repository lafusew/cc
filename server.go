package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lafusew/cc/controllers"
	"github.com/lafusew/cc/data"
)

var port = 8080

func main() {
	db := data.Connect(true)

	uController := controllers.UserController {
		Db: db,
	}

	http.HandleFunc("/users/", uController.HandleUsers)

	log.Printf("Server starting on port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
