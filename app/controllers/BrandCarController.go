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

// Create godoc
// @Summary Create a new brand car
// @Description Create a new brand car
// @Tags brand-cars
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param brand_car body models.BrandCar true "Brand Car object"
// @Success 201 {object} models.BrandCar
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/brand-cars [post]
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

// GetAll godoc
// @Summary Get all brand cars
// @Description Get a list of all brand cars
// @Tags brand-cars
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {array} models.BrandCar
// @Failure 500 {object} map[string]string
// @Router /api/cms/brand-cars [get]
func (bcc *BrandCarController) GetAll(c *gin.Context) {
	var brandCars []models.BrandCar
	if err := bcc.DB.Find(&brandCars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brandCars)
}

// GetByID godoc
// @Summary Get a brand car by ID
// @Description Get details of a specific brand car
// @Tags brand-cars
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Brand Car ID"
// @Success 200 {object} models.BrandCar
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/cms/brand-cars/{id} [get]
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

// Update godoc
// @Summary Update a brand car
// @Description Update details of a specific brand car
// @Tags brand-cars
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Brand Car ID"
// @Param brand_car body models.BrandCar true "Updated Brand Car object"
// @Success 200 {object} models.BrandCar
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/brand-cars/{id} [put]
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

// Delete godoc
// @Summary Delete a brand car
// @Description Delete a specific brand car
// @Tags brand-cars
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Brand Car ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/brand-cars/{id} [delete]
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
