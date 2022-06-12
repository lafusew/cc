package controllers

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (c *Controller) HandleTransactions(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		if parts[len(parts) - 1] == "" {
			c.GetAllTransactions(w, r)
			return
		} 
			
		c.GetTransaction(w, r)
		return
	case "POST":
		c.PostTransaction(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}