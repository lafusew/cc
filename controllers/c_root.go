package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lafusew/cc/data/models"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (c *Controller) JsonToModel(r *http.Request, w http.ResponseWriter, v interface{}) {
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

	err = json.Unmarshal(bodyBytes, &v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (c *Controller) JsonResponse(w http.ResponseWriter, v interface{}) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (c *Controller) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		parts := strings.Split(r.URL.Path, "/")
		if parts[len(parts)-1] == "" {
			users := c.GetAllUsers(0)
			c.JsonResponse(w, users)
			return
		}

		u := c.GetUserById(parts[2])
		c.JsonResponse(w, u)
		return
	case "POST":
		u := &models.User{}
		c.JsonToModel(r, w, u)
		u = c.PostUser(u)

		c.JsonResponse(w, u)
		return
	case "PUT":
		parts := strings.Split(r.URL.Path, "/")
		u := &models.User{}
		c.JsonToModel(r, w, u)

		u = c.PutUser(u, parts[2])
		c.JsonResponse(w, u)
		return
	case "DELETE":
		parts := strings.Split(r.URL.Path, "/")
		u := &models.User{}
		c.DeleteUser(u, parts[2])
		c.JsonResponse(w, map[string]interface{}{
			"deleted": true,
		})
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
