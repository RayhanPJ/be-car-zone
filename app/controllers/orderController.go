package controllers

import (
	"be-car-zone/app/models"
	"be-car-zone/app/pkg/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

// FindAll godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Order
// @Router /api/cms/orders [get]
func (ctrl *OrderController) FindAll(c *gin.Context) {
	var orders []models.Order
	if err := ctrl.DB.Preload("Car").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var orderDetails []models.OrderDetail
	for _, order := range orders {
		orderDetails = append(orderDetails, models.OrderDetail{
			ID:         order.ID,
			UserID:     order.UserID,
			CarID:      order.CarID,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
			Car: models.CarDetail{
				ID:          order.Car.ID,
				Name:        order.Car.Name,
				Description: order.Car.Description,
				ImageCar:    order.Car.ImageCar,
				Price:       order.Car.Price,
				TypeID:      order.Car.TypeID,
				BrandID:     order.Car.BrandID,
				IsSecond:    order.Car.IsSecond,
				CreatedAt:   order.Car.CreatedAt,
				UpdatedAt:   order.Car.UpdatedAt,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": orderDetails})
}

// FindByID godoc
// @Summary Get order by id
// @Description Get order by id
// @Tags orders
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Router /api/cms/orders/{id} [get]
func (ctrl *OrderController) FindByID(c *gin.Context) {
	var order models.Order
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Create godoc
// @Summary Create new order
// @Description Create new order
// @Tags orders
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param order body models.Order true "Order Data"
// @Success 200 {object} models.Order
// @Router /api/cms/orders [post]
func (ctrl *OrderController) Create(c *gin.Context) {
	var req models.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userId, _ = jwt.ExtractTokenID(c)

	// Cari user berdasarkan userID
	var user models.User
	if err := ctrl.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	newOrder := models.Order{
		UserID:     userId,
		CarID:      req.CarID,
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
		CreatedAt:  time.Now(),
	}

	if err := ctrl.DB.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newOrder})
}

// Update godoc
// @Summary Update order
// @Description Update order
// @Tags orders
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Order ID"
// @Param order body models.Order true "Order Data"
// @Success 200 {object} models.Order
// @Router /api/cms/orders/{id} [put]
func (ctrl *OrderController) Update(c *gin.Context) {
	var order models.Order
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	var userId, _ = jwt.ExtractTokenID(c)

	// Cari user berdasarkan userID
	var user models.User
	if err := ctrl.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var req models.Order
	// Update fields
	order.UserID = userId
	order.CarID = req.CarID
	order.TotalPrice = req.TotalPrice
	order.Status = req.Status
	order.UpdatedAt = time.Now()

	if err := ctrl.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Delete godoc
// @Summary Delete order
// @Description Delete order
// @Tags orders
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Router /api/cms/orders/{id} [delete]
func (ctrl *OrderController) Delete(c *gin.Context) {
	var order models.Order
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if err := ctrl.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})
}
