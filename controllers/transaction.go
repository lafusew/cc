package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Label  string `json:"label"`
	ID     string `json:"id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int `json:"amount"`
	Scale int `json:"scale"`
	Date   time.Time `json:"date"`
}

type TransactionsHandlers struct {
	sync.Mutex
	store map[string]Transaction
}

func (h *TransactionsHandlers) HandleTransactions(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		if parts[len(parts) - 1] == "" {
			h.GetAllTransactions(w, r)
			return
		} 
			
		h.GetTransaction(w, r)
		return
	case "POST":
		h.PostTransaction(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *TransactionsHandlers) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
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

func (h *TransactionsHandlers) GetTransaction(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	fmt.Sprintln(parts)
}

func (h *TransactionsHandlers) PostTransaction(w http.ResponseWriter, r *http.Request) {
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

	transaction.ID = uuid.New().String()
	transaction.Date = time.Now()

	log.Printf("New transaction registered with id: %s", transaction.ID)

	h.Lock()
	defer h.Unlock()

	h.store[transaction.ID] = transaction
}

func NewTransactionHandlers() *TransactionsHandlers {
	return &TransactionsHandlers{
		store: map[string]Transaction{
			"id1": {
				Label: "First transaction",
				ID: uuid.New().String(),
				From: uuid.New().String(),
				To: uuid.New().String(),
				Amount: 12,
				Scale: 0,
				Date: time.Now(),
			},
		},
	}
}