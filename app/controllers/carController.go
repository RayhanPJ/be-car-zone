package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"time"

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
	Sold        bool    `json:"sold"`
}

type CarController struct {
	DB *gorm.DB
}

type Result struct {
	Period string `json:"period"`
	Count  int    `json:"count"`
}

type WeeklyResult struct {
	Date   string `json:"date"`
	New    int    `json:"new"`
	Second int    `json:"second"`
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
		TypeID:      input.TypeID,
		BrandID:     input.BrandID,
		IsSecond:    input.IsSecond,
		Sold:        input.Sold,
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
	car.TypeID = input.TypeID
	car.BrandID = input.BrandID
	car.IsSecond = input.IsSecond
	car.Sold = input.Sold

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

// GetCarChartData handles requests for car sales data by week, month, and year
// @Summary Get car sales data
// @Description Get the number of cars sold per week, month, and per year
// @Tags Cars
// @Accept  json
// @Produce  json
// @Success 200 {object} gin.H{"weekly": []WeeklyResult, "monthly": []Result, "yearly": []Result}
// @Failure 500 {object} gin.H{"error": "Failed to get cars data"}
// @Router /cars/sales-data [get]
func (cc *CarController) GetCarChartData(c *gin.Context) {
	// Fetch sold cars data
	var cars []models.Car
	if err := cc.DB.Where("sold = ?", true).Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cars data"})
		return
	}

	// Calculate weekly, monthly, and yearly sales
	weeklySales := make(map[string]WeeklyResult)
	monthlySales := make(map[string]int)
	yearlySales := make(map[string]int)

	// Get the current time and calculate the time 7 days ago
	now := time.Now()
	weekAgo := now.AddDate(0, 0, -7)

	for _, car := range cars {
		createdDate := car.CreatedAt

		// Check if the car was sold in the last week
		if createdDate.After(weekAgo) && createdDate.Before(now) {
			date := createdDate.Format("2006-01-02")
			weeklyResult := weeklySales[date]
			if car.IsSecond {
				weeklyResult.Second++
			} else {
				weeklyResult.New++
			}
			weeklySales[date] = weeklyResult
		}

		// Calculate monthly and yearly sales
		month := createdDate.Format("2006-01")
		year := createdDate.Format("2006")
		monthlySales[month]++
		yearlySales[year]++
	}

	// Prepare results
	var weeklyResults []WeeklyResult
	for _, result := range weeklySales {
		weeklyResults = append(weeklyResults, result)
	}

	var monthlyResults, yearlyResults []Result
	for month, count := range monthlySales {
		monthlyResults = append(monthlyResults, Result{Period: month, Count: count})
	}
	for year, count := range yearlySales {
		yearlyResults = append(yearlyResults, Result{Period: year, Count: count})
	}

	// Sort results
	sort.Slice(weeklyResults, func(i, j int) bool {
		return weeklyResults[i].Date < weeklyResults[j].Date
	})
	sort.Slice(monthlyResults, func(i, j int) bool {
		return monthlyResults[i].Period < monthlyResults[j].Period
	})
	sort.Slice(yearlyResults, func(i, j int) bool {
		return yearlyResults[i].Period < yearlyResults[j].Period
	})

	// Respond with the data
	c.JSON(http.StatusOK, gin.H{"weekly": weeklyResults, "monthly": monthlyResults, "yearly": yearlyResults})
}
