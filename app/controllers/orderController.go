package controllers

import (
	"be-car-zone/app/models"
	"net/http"

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
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
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

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
