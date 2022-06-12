package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func (c *Controller) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := make([]Transaction, 10)

	jsonBytes, err := json.Marshal(transactions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (c *Controller) GetTransaction(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	fmt.Sprintln(parts)
}

func (c *Controller) PostTransaction(w http.ResponseWriter, r *http.Request) {
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
	log.Println(transaction)
}