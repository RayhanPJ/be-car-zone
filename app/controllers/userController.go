package controllers

import (
	"be-car-zone/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// FindAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.User
// @Router /api/cms/users [get]
func (ctrl *UserController) FindAll(c *gin.Context) {
	var users []models.User
	if err := ctrl.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var listUsers []models.UserList
	for _, res := range users {
		listUsers = append(listUsers, models.UserList{
			ID:       res.ID,
			Username: res.Username,
			Email:    res.Email,
			RoleName: res.Role.RoleName,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": listUsers})
}

// FindByID godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /api/cms/users/{id} [get]
func (ctrl *UserController) FindByID(c *gin.Context) {
	var user models.User
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Create godoc
// @Summary Create new user
// @Description Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User
// @Router /api/cms/users [post]
func (ctrl *UserController) Create(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Update godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "User ID"
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User
// @Router /api/cms/users/{id} [put]
func (ctrl *UserController) Update(c *gin.Context) {
	var user models.User
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Delete godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /api/cms/users/{id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {
	var user models.User
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if err := ctrl.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})
}
