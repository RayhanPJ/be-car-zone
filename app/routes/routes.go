package routes

import (
	"be-car-zone/app/controllers"
	"be-car-zone/app/middlewares"
	"be-car-zone/app/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Init controllers
	carController := &controllers.CarController{DB: db}
	brandCarController := &controllers.BrandCarController{DB: db}
	typeCarController := &controllers.TypeCarController{DB: db}
	orderController := &controllers.OrderController{DB: db}

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	authController := &controllers.AuthController{DB: db}
	userController := &controllers.UserController{DB: db}
	roleController := &controllers.RoleController{DB: db}
	transactionController := &controllers.TransactionController{DB: db}

	// Authentication User
	authRoute := r.Group("/api/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)
	authRoute.GET("/me", authController.GetCurrentUser, middlewares.JwtAuthMiddleware(utils.RoleUser, utils.RoleAdmin))
	authRoute.POST("/change-password", authController.ChangePassword, middlewares.JwtAuthMiddleware(utils.RoleUser, utils.RoleAdmin))

	// CMS Route
	cmsRoute := r.Group("/api/cms/", middlewares.JwtAuthMiddleware(utils.RoleAdmin))

	// CMS User
	cmsRoute.GET("/users", userController.FindAll)
	cmsRoute.GET("/users/:id", userController.FindByID)
	cmsRoute.POST("/users", userController.Create)
	cmsRoute.PUT("/users/:id", userController.Update)
	cmsRoute.DELETE("/users/:id", userController.Delete)

	// CMS Role
	cmsRoute.GET("/roles", roleController.FindAll)
	cmsRoute.GET("/roles/:id", roleController.FindByID)
	cmsRoute.POST("/roles", roleController.Create)
	cmsRoute.PUT("/roles/:id", roleController.Update)
	cmsRoute.DELETE("/roles/:id", roleController.Delete)

	// CMS Order
	cmsRoute.GET("/orders", orderController.FindAll)
	cmsRoute.GET("/orders/:id", orderController.FindByID)
	cmsRoute.POST("/orders", orderController.Create)
	cmsRoute.PUT("/orders/:id", orderController.Update)
	cmsRoute.DELETE("/orders/:id", orderController.Delete)

	// CMS Transaction
	cmsRoute.GET("/transactions", transactionController.FindAll)
	cmsRoute.GET("/transactions/:id", transactionController.FindByID)
	cmsRoute.POST("/transactions", transactionController.Create)
	cmsRoute.PUT("/transactions/:id", transactionController.Update)
	cmsRoute.DELETE("/transactions/:id", transactionController.Delete)

	// Car
	cmsRoute.POST("/cars", carController.Create)
	cmsRoute.GET("/cars", carController.GetAll)
	cmsRoute.GET("/cars/:id", carController.GetByID)
	cmsRoute.PUT("/cars/:id", carController.Update)
	cmsRoute.DELETE("/cars/:id", carController.Delete)

	// BrandCar
	cmsRoute.POST("/brand-cars", brandCarController.Create)
	cmsRoute.GET("/brand-cars", brandCarController.GetAll)
	cmsRoute.GET("/brand-cars/:id", brandCarController.GetByID)
	cmsRoute.PUT("/brand-cars/:id", brandCarController.Update)
	cmsRoute.DELETE("/brand-cars/:id", brandCarController.Delete)

	// TypeCar
	cmsRoute.POST("/type-cars", typeCarController.Create)
	cmsRoute.GET("/type-cars", typeCarController.GetAll)
	cmsRoute.GET("/type-cars/:id", typeCarController.GetByID)
	cmsRoute.PUT("/type-cars/:id", typeCarController.Update)
	cmsRoute.DELETE("/type-cars/:id", typeCarController.Delete)

}
