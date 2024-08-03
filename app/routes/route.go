package routes

import (
	"be-car-zone/app/controllers"
	"be-car-zone/app/middlewares"
	"be-car-zone/app/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, r *gin.Engine) {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	authController := &controllers.AuthController{DB: db}
	userController := &controllers.UserController{DB: db}
	roleController := &controllers.RoleController{DB: db}

	// Authentication User
	authRoute := r.Group("/api/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)
	authRoute.GET("/me", authController.GetCurrentUser, middlewares.JwtAuthMiddleware(utils.RoleUser, utils.RoleAdmin))

	// CMS Routes
	cmsRoute := r.Group("/api/cms", middlewares.JwtAuthMiddleware(utils.RoleAdmin))
	// Manage Users
	cmsRoute.GET("/users", userController.FindAll)
	cmsRoute.GET("/users/:id", userController.FindByID)
	cmsRoute.POST("/users", userController.Create)
	cmsRoute.PUT("/users/:id", userController.Update)
	cmsRoute.DELETE("/users/:id", userController.Delete)
	// Manage Roles
	cmsRoute.GET("/roles", roleController.FindAll)
	cmsRoute.GET("/roles/:id", roleController.FindByID)
	cmsRoute.POST("/roles", roleController.Create)
	cmsRoute.PUT("/roles:id", roleController.Update)
	cmsRoute.DELETE("/roles/:id", roleController.Delete)

}
