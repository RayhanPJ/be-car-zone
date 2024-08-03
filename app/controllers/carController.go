package controllers

import (
	"be-car-zone/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarController struct {
	DB *gorm.DB
}

func (cc *CarController) Create(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, car)
}

func (cc *CarController) GetAll(c *gin.Context) {
	var cars []models.Car
	if err := cc.DB.Preload("Type").Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (cc *CarController) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var car models.Car
	if err := cc.DB.Preload("Type").First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, car)
}

func (cc *CarController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var car models.Car
	if err := cc.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

func (cc *CarController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := cc.DB.Delete(&models.Car{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
