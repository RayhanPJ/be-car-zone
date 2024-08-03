package controllers

import (
	"be-car-zone/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TypeCarController struct {
	DB *gorm.DB
}

func (tcc *TypeCarController) Create(c *gin.Context) {
	var typeCar models.TypeCar
	if err := c.ShouldBindJSON(&typeCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tcc.DB.Create(&typeCar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, typeCar)
}

func (tcc *TypeCarController) GetAll(c *gin.Context) {
	var typeCars []models.TypeCar
	if err := tcc.DB.Find(&typeCars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, typeCars)
}

func (tcc *TypeCarController) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var typeCar models.TypeCar
	if err := tcc.DB.Preload("Cars").First(&typeCar, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type car not found"})
		return
	}

	c.JSON(http.StatusOK, typeCar)
}

func (tcc *TypeCarController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var typeCar models.TypeCar
	if err := tcc.DB.First(&typeCar, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type car not found"})
		return
	}

	if err := c.ShouldBindJSON(&typeCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tcc.DB.Save(&typeCar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, typeCar)
}

func (tcc *TypeCarController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := tcc.DB.Delete(&models.TypeCar{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Type car deleted successfully"})
}
