package controllers

import (
	"be-car-zone/app/models"
	"be-car-zone/app/pkg/jwt"
	"be-car-zone/app/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body models.LoginRequest true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := utils.NewValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user *models.User
	if err := ctrl.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	if user == nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := jwt.GenerateToken(user.ID, uint(user.RoleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body models.RegisterRequest true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := utils.NewValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := ctrl.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		if existingUser.Username == req.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
			return
		}
		if existingUser.Email == req.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID:   utils.IDRoleUser,
	}

	if err := ctrl.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser.Username})
}

// GetCurrentUser godoc
// @Summary Get Current User by token.
// @Description Get Current User by token.
// @Tags Auth
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.User
// @Router /api/auth/me [get]
func (ctrl *AuthController) GetCurrentUser(c *gin.Context) {

	var userId, _ = jwt.ExtractTokenID(c)

	var user models.User

	if err := ctrl.DB.Where("id = ?", userId).Preload("Role").First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// ChangePasswordUser godoc
// @Summary Change Password User by token.
// @Description Change Password User by token.
// @Tags Auth
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param body body models.InputChangePassword true "Change Password Request Body"
// @Security BearerToken
// @Success 200 {object} map[string]interface{} "{"message": "Password changed successfully"}"
// @Failure 400 {object} map[string]interface{} "{"error": "string"}"
// @Failure 401 {object} map[string]interface{} "{"error": "string"}"
// @Failure 500 {object} map[string]interface{} "{"error": "string"}"
// @Router /api/auth/change-password [post]
func (ctrl *AuthController) ChangePassword(c *gin.Context) {
	var input models.InputChangePassword
	if err := c.ShouldBindJSON(&input); err != nil {
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

	// Verifikasi old password
	if !utils.CheckPasswordHash(input.OldPassword, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
		return
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update password user
	user.Password = hashedPassword
	if err := ctrl.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
