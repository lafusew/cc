package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lafusew/cc/controllers"
)

var port = 8080;

func main() {
	transactionsHandlers := controllers.NewTransactionHandlers()

	http.HandleFunc("/transactions/", transactionsHandlers.HandleTransactions)
	http.HandleFunc("/transactions", transactionsHandlers.HandleTransactions)

	log.Printf("Server starting on port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
