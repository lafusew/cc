package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/lafusew/cc/data/models"
	"github.com/lafusew/cc/res"
	"gorm.io/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func (c *UserController) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		parts := strings.Split(r.URL.Path, "/")
		if parts[len(parts)-1] == "" {
			users := c.GetAllUsers(0)
			res.JsonResponse(w, users)
			return
		}

		u := c.GetUserById(parts[2])
		res.JsonResponse(w, u)
		return
	case "POST":
		u := &models.User{}
		res.JsonToModel(r, w, u)
		u = c.PostUser(u)

		res.JsonResponse(w, u)
		return
	case "PUT":
		parts := strings.Split(r.URL.Path, "/")
		u := &models.User{}
		res.JsonToModel(r, w, u)

		u = c.PutUser(u, parts[2])
		res.JsonResponse(w, u)
		return
	case "DELETE":
		parts := strings.Split(r.URL.Path, "/")
		u := &models.User{}
		err := c.DeleteUser(u, parts[2])
		if err != nil {
			res.JsonResponse(w, map[string]interface{}{
				"id": parts[2],
				"error": err.Error(),
			})			
			return
		} 

		res.JsonResponse(w, map[string]interface{}{
			"id": parts[2],
			"deleted": true,
		})
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}


func (c *UserController) GetUserById(idString string) *models.User {
	id, err := uuid.Parse(idString)
	if err != nil {
		log.Println(err)
		return nil
	}
	var u = &models.User{ID: id}

	u, err = u.FindById(c.Db, u.ID)
	if err != nil {
		log.Println(err)
		return nil
	}

	return u
}

func (c *UserController) GetAllUsers(pagination int) *[]models.User {
	u := &models.User{}
	users, err := u.FindAll(c.Db, pagination, 100)
	if err != nil {
		log.Println(err)
		return nil
	}

	return users
}

func (c *UserController) PostUser(u *models.User) *models.User {
	u, err := u.Create(c.Db)
	if err != nil {
		log.Printf("error while saving user: %s", err.Error())
		return nil
	}

	return u
}

func (c *UserController) PutUser(u *models.User, idString string) *models.User {
	id, err := uuid.Parse(idString)
	if err != nil {
		log.Println(err)
		return nil
	}
	u.Update(c.Db, id)

	return u
}

func (c *UserController) DeleteUser(u *models.User, idString string) error {
	id, err := uuid.Parse(idString)
	if err != nil {
		log.Println(err)
		return err
	}

	err = u.Delete(c.Db, id)

	return err
}
