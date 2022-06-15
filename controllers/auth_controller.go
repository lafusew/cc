package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lafusew/cc/data/models"
	"github.com/lafusew/cc/utils"
	"gorm.io/gorm"
)

type SignupPayload struct {
	DisplayName string `json:"display_name"`
	Tag string `json:"tag"`
	Identifier string `json:"identifier"`
	Password string `json:"password"`
}

type SigninPayload struct {
	Identifier string `json:"identifier"`
	Password string `json:"password"`
}

type AuthController struct {
	Db *gorm.DB
}

func (c *AuthController) HandleAuth(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/");
	
	switch r.Method {
	case "POST":
		if parts[1] == "signin" {
			p := &SigninPayload{}
			utils.JsonToModel(r, w, p)

			a := &models.Auth{}
			a.FindByIdentifier(c.Db, p.Identifier)

			if (a.Identifier != p.Identifier || a.Password != p.Password) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Wrong password or identifier"))
				return
			}

			token, err := utils.IssueJWT(a.UserID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Something went wrong while generating your credentials token: %s", err.Error())))
				return
			}

			utils.JsonResponse(w, token)
			return
		}

		if parts[1] == "signup" {
			p := &SignupPayload{}
			utils.JsonToModel(r, w, p)


			u := &models.User{DisplayName: p.DisplayName, Tag: p.Tag}
			a := &models.Auth{Identifier: p.Identifier, Password: p.Password }


			err := u.Create(c.Db)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Something went wrong while creating your account: %s", err.Error())))
				return
			}

			err = a.Create(c.Db, u.ID)
			if err != nil {
				u.Delete(c.Db, u.ID)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Something went wrong while creating your account: %s", err.Error())))
				return
			}

			token, err := utils.IssueJWT(u.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Something went wrong while generating your credentials token: %s", err.Error())))
				return
			}

			utils.JsonResponse(w, token)
		}
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

}