package controllers

import (
	"be-car-zone/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleController struct {
	DB *gorm.DB
}

// FindAll godoc
// @Summary Get all roles for Admin
// @Description Get all roles for Admin
// @Tags roles
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Role
// @Router /api/cms/roles [get]
func (ctrl *RoleController) FindAll(c *gin.Context) {
	var roles []models.Role
	if err := ctrl.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var listRoles []models.RoleList
	for _, res := range roles {
		listRoles = append(listRoles, models.RoleList{
			ID:       res.ID,
			RoleName: res.RoleName,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": listRoles})
}

// FindByID godoc
// @Summary Get role by id
// @Description Get role by id
// @Tags roles
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Role ID"
// @Success 200 {object} models.Role
// @Router /roles/{id} [get]
func (ctrl *RoleController) FindByID(c *gin.Context) {
	var role models.Role
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// Create godoc
// @Summary Create new role
// @Description Create new role
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param role body models.Role true "Role Data"
// @Success 200 {object} models.Role
// @Router /api/cms/roles [post]
func (ctrl *RoleController) Create(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := ctrl.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error when creating role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// Update godoc
// @Summary Update role
// @Description Update role
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Role ID"
// @Param role body models.RoleRequest true "Role Data"
// @Success 200 {object} models.Role
// @Router /api/cms/roles/{id} [put]
func (ctrl *RoleController) Update(c *gin.Context) {
	var role models.Role

	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := ctrl.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error when updating role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// Delete godoc
// @Summary Delete role
// @Description Delete role
// @Tags roles
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Role ID"
// @Success 200 {object} models.Role
// @Router /api/cms/roles/{id} [delete]
func (ctrl *RoleController) Delete(c *gin.Context) {
	var role models.Role
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if err := ctrl.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error when deleting role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}
