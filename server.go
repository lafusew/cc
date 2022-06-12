package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lafusew/cc/controllers"
	"github.com/lafusew/cc/data"
)

var c = controllers.Controller{}
var port = 8080

func main() {
	c.Db = data.Connect()
	data.Init(c.Db)

	http.HandleFunc("/transactions/", c.HandleTransactions)

	log.Printf("Server starting on port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
