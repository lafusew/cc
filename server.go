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
	c := controllers.Controller{
		Db: data.Connect(true),
	}

	http.HandleFunc("/users/", c.HandleUsers)

	log.Printf("Server starting on port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
