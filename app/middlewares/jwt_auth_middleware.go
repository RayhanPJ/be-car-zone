package middlewares

import (
	"be-car-zone/app/models"
	"be-car-zone/app/pkg/jwt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// JwtAuthMiddleware checks the validity of the JWT and authorizes based on user role
func JwtAuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userId, err := jwt.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		var user models.User
		db := c.MustGet("db").(*gorm.DB)
		findUserErr := db.Preload("Role").Where("id = ?", userId).First(&user).Error
		if findUserErr != nil {
			c.String(http.StatusUnauthorized, findUserErr.Error())
			c.Abort()
			return
		}

		// Check if the user's role is allowed to access the route
		for _, role := range allowedRoles {
			if role == user.Role.RoleName || user.Role.RoleName == "admin" {
				c.Set("user_id", user.ID)
				c.Set("user_role", user.Role.RoleName)
				c.Next()
				return
			}
		}

		roleError := errors.New("sorry, your role cannot access this route")
		c.String(http.StatusForbidden, roleError.Error())
		c.Abort()
	}
}
