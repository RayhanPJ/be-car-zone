package controllers

import (
	"net/http"
	"strconv"

	"be-car-zone/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	ImageCar    string  `json:"image_car"`
	Price       float64 `json:"price" binding:"required"`
	ImageUrl    string  `json:"image_url"`
	TypeID      uint    `json:"type_id" binding:"required"`
	BrandID     uint    `json:"brand_id" binding:"required"`
	IsSecond    bool    `json:"is_second"`
}

type CarController struct {
	DB *gorm.DB
}

// Create godoc
// @Summary Create a new car
// @Description Create a new car
// @Tags cars
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param car body CarInput true "Car object"
// @Success 201 {object} models.Car
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/cars [post]
func (cc *CarController) Create(c *gin.Context) {
	var input CarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := models.Car{
		Name:        input.Name,
		Description: input.Description,
		ImageCar:    input.ImageCar,
		Price:       input.Price,
		ImageUrl:    input.ImageUrl,
		TypeID:      input.TypeID,
		BrandID:     input.BrandID,
		IsSecond:    input.IsSecond,
	}

	if err := cc.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	if err := cc.DB.Preload("Type").Preload("Brand").First(&car, car.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load car details"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Car created successfully", "car": car})
}

// GetAll godoc
// @Summary Get all cars
// @Description Get a list of all cars with their types and brands
// @Tags cars
// @Produce json
// @Success 200 {array} models.Car
// @Failure 500 {object} map[string]string
// @Router /api/cms/cars [get]
func (cc *CarController) GetAll(c *gin.Context) {
	var cars []models.Car
	if err := cc.DB.Preload("Type").Preload("Brand").Order("created_at DESC").Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cars"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

// GetByID godoc
// @Summary Get a car by ID
// @Description Get details of a specific car including its type and brand
// @Tags cars
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} models.Car
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/cms/cars/{id} [get]
func (cc *CarController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var car models.Car
	if err := cc.DB.Preload("Type").Preload("Brand").First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"car": car})
}

// Update godoc
// @Summary Update a car
// @Description Update details of a specific car
// @Tags cars
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Car ID"
// @Param car body CarInput true "Updated Car object"
// @Success 200 {object} models.Car
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/cars/{id} [put]
func (cc *CarController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var car models.Car
	if err := cc.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	var input CarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car.Name = input.Name
	car.Description = input.Description
	car.Price = input.Price
	car.ImageUrl = input.ImageUrl
	car.TypeID = input.TypeID
	car.BrandID = input.BrandID
	car.IsSecond = input.IsSecond

	if err := cc.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}

	if err := cc.DB.Preload("Type").Preload("Brand").First(&car, car.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load car details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully", "car": car})
}

// Delete godoc
// @Summary Delete a car
// @Description Delete a specific car
// @Tags cars
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Car ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cms/cars/{id} [delete]
func (cc *CarController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var car models.Car
	if err := cc.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	if err := cc.DB.Delete(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
