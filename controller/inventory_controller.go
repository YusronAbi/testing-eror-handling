package controllers

import (
	"inventory-management/database"
	"inventory-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
	var inventory []models.Inventory
	database.DB.Find(&inventory)
	c.JSON(http.StatusOK, inventory)
}

func UpdateInventory(c *gin.Context) {
	var inventory models.Inventory
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	var input struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inventory.Quantity = input.Quantity
	database.DB.Save(&inventory)
	c.JSON(http.StatusOK, inventory)
}
func CreateInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&inventory)
	c.JSON(http.StatusCreated, inventory)
}
