package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Transaction struct {
	Label  string `json:"label"`
	ID     string `json:"id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int `json:"amount"`
	Date   time.Time `json:"date"`
}

type transactionHandlers struct {
	sync.Mutex
	store map[string]Transaction
}

func (h *transactionHandlers) transactions(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s, transactions controller \n",r.Method)
	switch r.Method {
	case "GET":
		h.getAll(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *transactionHandlers) transaction(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s, transaction controller \n", r.Method)
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *transactionHandlers) get(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	fmt.Sprintln(parts)

	// h.Lock()
	// i := 0
	// for _, transaction := range h.store {
	// 	transactions[i] = transaction;
	// 	i++
	// }
	// h.Unlock()

	// jsonBytes, err := json.Marshal(transactions)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }

	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(jsonBytes)
}

func (h *transactionHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var transaction Transaction
	err = json.Unmarshal(bodyBytes, &transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	transaction.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	fmt.Println("ok")

	h.Lock()
	defer h.Unlock()

	h.store[transaction.ID] = transaction
}

func (h *transactionHandlers) getAll(w http.ResponseWriter, r *http.Request) {
	transactions := make([]Transaction, len(h.store))

	h.Lock()
	i := 0
	for _, transaction := range h.store {
		transactions[i] = transaction;
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(transactions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newTransactionHandlers() *transactionHandlers {
	return &transactionHandlers{
		store: map[string]Transaction{
			"id1": {
				Label: "First transaction",
				ID: "iojd23",
				From: "dei12",
				To: "d234d",
				Amount: 12,
				Date: time.Now(),
			},
		},
	}
}