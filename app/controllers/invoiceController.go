package controllers

import (
	"be-car-zone/app/models"
	"be-car-zone/app/pkg/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InvoiceController struct {
	DB *gorm.DB
}

// FindAll godoc
// @Summary Get all invoices
// @Description Get all invoices
// @Tags invoices
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Invoice
// @Router /api/cms/invoices [get]
func (ctrl *InvoiceController) FindAll(c *gin.Context) {
	var invoices []models.Invoice
	if err := ctrl.DB.Preload("Order").Preload("Transaction").Order("created_at DESC").Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var invoiceDetails []models.InvoiceDetail
	for _, invoice := range invoices {
		// Filter untuk hanya memproses invoices dengan status true
		if invoice.Order.Status == true {
			invoiceDetails = append(invoiceDetails, models.InvoiceDetail{
				ID:            invoice.ID,
				OrderID:       invoice.OrderID,
				TransactionID: invoice.TransactionID,
				CreatedAt:     invoice.CreatedAt,
				UpdatedAt:     invoice.UpdatedAt,
				Order: models.OrderDetail{
					ID:         invoice.Order.ID,
					UserID:     invoice.Order.UserID,
					CarID:      invoice.Order.CarID,
					TotalPrice: invoice.Order.TotalPrice,
					Status:     invoice.Order.Status,
					OrderImage: invoice.Order.OrderImage,
					CreatedAt:  invoice.Order.CreatedAt,
					UpdatedAt:  invoice.Order.UpdatedAt,
				},
				Transaction: models.TransactionDetail{
					ID:               invoice.Transaction.ID,
					OrderID:          invoice.Transaction.OrderID,
					PaymentProvider:  invoice.Transaction.PaymentProvider,
					NoRek:            invoice.Transaction.NoRek,
					Amount:           invoice.Transaction.Amount,
					TransactionDate:  invoice.Transaction.TransactionDate,
					CreatedAt:        invoice.Transaction.CreatedAt,
					UpdatedAt:        invoice.Transaction.UpdatedAt,
				},
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": invoiceDetails})
}


// FindByID godoc
// @Summary Get invoice by id
// @Description Get invoice by id
// @Tags invoices
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "User ID"
// @Success 200 {object} models.Invoice
// @Router /api/cms/invoices/{id} [get]
func (ctrl *InvoiceController) FindByID(c *gin.Context) {
	var invoice []models.Invoice
	if err := ctrl.DB.Where("user_id = ?", c.Param("id")).Find(&invoice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

// Create godoc
// @Summary Create new invoice
// @Description Create new invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param invoice body models.Invoice true "Invoice Data"
// @Success 200 {object} models.Invoice
// @Router /api/cms/invoices [post]
func (ctrl *InvoiceController) Create(c *gin.Context) {
	var req models.Invoice
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

	newInvoice := models.Invoice{
		OrderID:       req.OrderID,
		TransactionID: req.TransactionID,
		CreatedAt:     time.Now(),
	}

	if err := ctrl.DB.Create(&newInvoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newInvoice})
}

// Update godoc
// @Summary Update invoice
// @Description Update invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Invoice ID"
// @Param invoice body models.Invoice true "Invoice Data"
// @Success 200 {object} models.Invoice
// @Router /api/cms/invoices/{id} [put]
func (ctrl *InvoiceController) Update(c *gin.Context) {
	var invoice models.Invoice
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&invoice).Error; err != nil {
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

	var req models.Invoice
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update fields
	invoice.OrderID = req.OrderID
	invoice.TransactionID = req.TransactionID
	invoice.UpdatedAt = time.Now()

	if err := ctrl.DB.Save(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

// Delete godoc
// @Summary Delete invoice
// @Description Delete invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Invoice ID"
// @Success 200 {object} models.Invoice
// @Router /api/cms/invoices/{id} [delete]
func (ctrl *InvoiceController) Delete(c *gin.Context) {
	var invoice models.Invoice
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&invoice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if err := ctrl.DB.Delete(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})
}
