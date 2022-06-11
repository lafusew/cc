package main

import (
	"net/http"
)

func main() {
	transactionsHandlers := newTransactionHandlers()

	http.HandleFunc("/transactions", transactionsHandlers.transactions)
	http.HandleFunc("/transaction/", transactionsHandlers.transaction)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
