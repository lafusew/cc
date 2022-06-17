package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/lafusew/cc/data/models"
	"github.com/lafusew/cc/utils"
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
			users, err := c.GetAllUsers(0)
			if err != nil {
				utils.JsonResponse(w, http.StatusNotFound,err.Error())
			}
			
			utils.JsonResponse(w, http.StatusOK, users)
			return
		}
		
		u := &models.User{}
		c.GetUserById(u ,parts[2])

		utils.JsonResponse(w, http.StatusOK, u)
		return
	case "POST":
		u := &models.User{}
		utils.JsonToModel(r, w, u)

		c.PostUser(u)

		utils.JsonResponse(w, http.StatusOK, u)
		return
	case "PUT":
		parts := strings.Split(r.URL.Path, "/")
		u := &models.User{}
		utils.JsonToModel(r, w, u)

		c.PutUser(u, parts[2])

		utils.JsonResponse(w, http.StatusOK, u)
		return
	case "DELETE":
		parts := strings.Split(r.URL.Path, "/")
		u := &models.User{}
		err := c.DeleteUser(u, parts[2])
		if err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"id": parts[2],
				"deleted": false,
				"error": err.Error(),
			})
			return
		} 

		utils.JsonResponse(w, http.StatusOK, map[string]interface{}{
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


func (c *UserController) GetUserById(u *models.User, idString string) error {
	id, err := uuid.Parse(idString)
	if err != nil {
		log.Println(err)
		return err
	}
	
	u.ID = id

	err = u.FindById(c.Db, u.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func (c *UserController) GetAllUsers(pagination int) (*[]models.User, error) {
	u := &models.User{}
	us, err := u.FindAll(c.Db, pagination, 100)
	if err != nil {
		return nil, err
	}

	return us, err
}

func (c *UserController) PostUser(u *models.User) error {
	err := u.Create(c.Db)
	if err != nil {
		log.Printf("error while saving user: %s", err.Error())
		return err
	}

	return err
}

func (c *UserController) PutUser(u *models.User, idString string) error {
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}
	u.Update(c.Db, id)

	return err
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
