package controllers

import (
	"be-car-zone/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BrandCarController struct {
	DB *gorm.DB
}

func (bcc *BrandCarController) Create(c *gin.Context) {
	var brandCar models.BrandCar
	if err := c.ShouldBindJSON(&brandCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bcc.DB.Create(&brandCar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brandCar)
}

func (bcc *BrandCarController) GetAll(c *gin.Context) {
	var brandCars []models.BrandCar
	if err := bcc.DB.Find(&brandCars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brandCars)
}

func (bcc *BrandCarController) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var brandCar models.BrandCar
	if err := bcc.DB.First(&brandCar, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand car not found"})
		return
	}

	c.JSON(http.StatusOK, brandCar)
}

func (bcc *BrandCarController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var brandCar models.BrandCar
	if err := bcc.DB.First(&brandCar, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand car not found"})
		return
	}

	if err := c.ShouldBindJSON(&brandCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bcc.DB.Save(&brandCar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brandCar)
}

func (bcc *BrandCarController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := bcc.DB.Delete(&models.BrandCar{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand car deleted successfully"})
}
