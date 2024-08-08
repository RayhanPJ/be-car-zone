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

// Create godoc
// @Summary Create a new type car
// @Description Create a new type car
// @Tags type-cars
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param type_car body models.TypeCar true "Type Car object"
// @Success 201 {object} models.TypeCar
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/type-cars [post]
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

// GetAll godoc
// @Summary Get all type cars
// @Description Get a list of all type cars
// @Tags type-cars
// @Produce json
// @Success 200 {array} models.TypeCar
// @Failure 500 {object} map[string]string
// @Router /api/cms/type-cars [get]
func (tcc *TypeCarController) GetAll(c *gin.Context) {
	var typeCars []models.TypeCar
	if err := tcc.DB.Order("created_at DESC").Find(&typeCars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, typeCars)
}

// GetByID godoc
// @Summary Get a type car by ID
// @Description Get details of a specific type car including associated cars
// @Tags type-cars
// @Produce json
// @Param id path int true "Type Car ID"
// @Success 200 {object} models.TypeCar
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/cms/type-cars/{id} [get]
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

// Update godoc
// @Summary Update a type car
// @Description Update details of a specific type car
// @Tags type-cars
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Type Car ID"
// @Param type_car body models.TypeCar true "Updated Type Car object"
// @Success 200 {object} models.TypeCar
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/type-cars/{id} [put]
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

// Delete godoc
// @Summary Delete a type car
// @Description Delete a specific type car
// @Tags type-cars
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Type Car ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/type-cars/{id} [delete]
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
