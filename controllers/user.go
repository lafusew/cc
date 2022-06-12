package controllers

import (
	"log"

	"github.com/google/uuid"
	"github.com/lafusew/cc/data/models"
)

func (c *Controller) GetUserById(idString string) *models.User {
	id, err := uuid.Parse(idString)
	if err != nil {
		log.Println(err)
		return nil
	}
	var u = &models.User{ ID: id }

	u, err = u.FindById(c.Db, u.ID)
	if err != nil {
		log.Println(err)
		return nil
	}

	return u
}

func (c *Controller) GetAllUsers(pagination int) *[]models.User {
	u := &models.User{}
	users, err := u.FindAll(c.Db, pagination, 100)
	if err != nil {
		log.Println(err)
		return nil
	}

	return users
}

func (c *Controller) PostUser(u *models.User) *models.User {
	u, err := u.Create(c.Db)
	if err != nil {
		log.Printf("error while saving user: %s", err.Error())
		return nil
	}

	return u
}

func (c *Controller) PutUser(u *models.User, idString string) *models.User {
	id, err := uuid.Parse(idString)
	if err != nil {
		log.Println(err)
		return nil
	}
	u.Update(c.Db, id)

	return u
}