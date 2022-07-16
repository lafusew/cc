package controllers

import (
	"net/http"
	"strings"

	"github.com/lafusew/cc/data/models"
	"github.com/lafusew/cc/utils"
	"gorm.io/gorm"
)

type CoinController struct {
	Db *gorm.DB
}

func (c *CoinController) HandleCoins(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		parts := strings.Split(r.URL.Path, "/")
		if parts[len(parts)-1] == "" {
			coins, err := c.GetAllCoins(0)
			if err != nil {
				utils.JsonResponse(w, http.StatusNotFound,err.Error())
			}
			utils.JsonResponse(w, http.StatusOK, coins)
			return
		}
		coin := &models.Coin{}
		c.GetCoinById(coin ,parts[2])
		utils.JsonResponse(w, http.StatusOK, coin)
		return
	case "POST":
		coin := &models.Coin{}
		utils.JsonToModel(r, w, coin)
		c.PostCoin(coin)
		utils.JsonResponse(w, http.StatusOK, coin)
		return
	case "PUT":
		parts := strings.Split(r.URL.Path, "/")
		coin := &models.Coin{}
		utils.JsonToModel(r, w, coin)
		c.PutCoin(coin, parts[2])
		utils.JsonResponse(w, http.StatusOK, coin)
		return
	case "DELETE":
		parts := strings.Split(r.URL.Path, "/")
		coin := &models.Coin{}
		err := c.DeleteCoin(coin, parts[2])
		if err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		utils.JsonResponse(w, http.StatusOK, map[string]interface{}{
			"message": "Coin deleted",
		})
		return
	}
}


func (c *CoinController) GetAllCoins(page int) ([]models.Coin, error) {
	coins := []models.Coin{}
	err := c.Db.Find(&coins).Error
	if err != nil {
		return nil, err
	}
	return coins, nil
}

func (c *CoinController) GetCoinById(coin *models.Coin, id string) error {
	err := c.Db.First(coin, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CoinController) PostCoin(coin *models.Coin) error {
	err := c.Db.Create(coin).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CoinController) PutCoin(coin *models.Coin, id string) error {
	err := c.Db.Save(coin).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CoinController) DeleteCoin(coin *models.Coin, id string) error {
	err := c.Db.Delete(coin, id).Error
	if err != nil {
		return err
	}
	return nil
}
