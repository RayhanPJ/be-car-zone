package controllers

import (
	"be-car-zone/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionController struct {
	DB *gorm.DB
}

// FindAll godoc
// @Summary Get all transactions
// @Description Get all transactions
// @Tags transactions
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Transaction
// @Router /api/cms/transactions [get]
func (ctrl *TransactionController) FindAll(c *gin.Context) {
	var transactions []models.Transaction
	if err := ctrl.DB.Preload("Order.Car").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var transactionDetails []models.TransactionDetail
	for _, transaction := range transactions {
		transactionDetails = append(transactionDetails, models.TransactionDetail{
			ID:               transaction.ID,
			OrderID:          transaction.OrderID,
			PaymentProvider:  transaction.PaymentProvider,
			TransactionImage: transaction.TransactionImage,
			NoRek:            transaction.NoRek,
			Amount:           transaction.Amount,
			TransactionDate:  transaction.TransactionDate,
			CreatedAt:        transaction.CreatedAt,
			UpdatedAt:        transaction.UpdatedAt,
			Order: models.OrderDetail{
				ID:         transaction.Order.ID,
				UserID:     transaction.Order.UserID,
				CarID:      transaction.Order.CarID,
				TotalPrice: transaction.Order.TotalPrice,
				Status:     transaction.Order.Status,
				CreatedAt:  transaction.Order.CreatedAt,
				UpdatedAt:  transaction.Order.UpdatedAt,
				Car: models.CarDetail{
					ID:          transaction.Order.Car.ID,
					Name:        transaction.Order.Car.Name,
					Description: transaction.Order.Car.Description,
					Price:       transaction.Order.Car.Price,
					TypeID:      transaction.Order.Car.TypeID,
					BrandID:     transaction.Order.Car.BrandID,
					IsSecond:    transaction.Order.Car.IsSecond,
					CreatedAt:   transaction.Order.Car.CreatedAt,
					UpdatedAt:   transaction.Order.Car.UpdatedAt,
				},
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": transactionDetails})
}

// FindByID godoc
// @Summary Get transaction by id
// @Description Get transaction by id
// @Tags transactions
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Router /api/cms/transactions/{id} [get]
func (ctrl *TransactionController) FindByID(c *gin.Context) {
	var transaction models.Transaction
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// Create godoc
// @Summary Create new transaction
// @Description Create new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param transaction body models.Transaction true "Transaction Data"
// @Success 200 {object} models.Transaction
// @Router /api/cms/transactions [post]
func (ctrl *TransactionController) Create(c *gin.Context) {
	var req models.Transaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction := models.Transaction{
		OrderID:          req.OrderID,
		PaymentProvider:  req.PaymentProvider,
		TransactionImage: req.TransactionImage,
		NoRek:            req.NoRek,
		Amount:           req.Amount,
		TransactionDate:  time.Now(),
		CreatedAt:        time.Now(),
	}

	if err := ctrl.DB.Create(&newTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newTransaction})
}

// Update godoc
// @Summary Update transaction
// @Description Update transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Transaction ID"
// @Param transaction body models.Transaction true "Transaction Data"
// @Success 200 {object} models.Transaction
// @Router /api/cms/transactions/{id} [put]
func (ctrl *TransactionController) Update(c *gin.Context) {
	var transaction models.Transaction
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	var req models.Transaction
	// Update fields
	transaction.OrderID = req.OrderID
	transaction.PaymentProvider = req.PaymentProvider
	transaction.TransactionImage = req.TransactionImage
	transaction.NoRek = req.NoRek
	transaction.Amount = req.Amount
	transaction.UpdatedAt = time.Now()

	if err := ctrl.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// Delete godoc
// @Summary Delete transaction
// @Description Delete transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Router /api/cms/transactions/{id} [delete]
func (ctrl *TransactionController) Delete(c *gin.Context) {
	var transaction models.Transaction
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if err := ctrl.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})
}
