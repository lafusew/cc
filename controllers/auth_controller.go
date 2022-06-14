package controllers

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type AuthController struct {
	Db *gorm.DB
}

func (c *AuthController) HandleAuth(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/");
	
	switch r.Method {
	case "POST":
		if parts[1] == "signin" {
			
		}

		if parts[1] == "signup" {

		}
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

}