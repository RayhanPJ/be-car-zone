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
	cmsRouteAdmin := r.Group("/api/cms/", middlewares.JwtAuthMiddleware(utils.RoleAdmin))
	cmsRouteAllRole := r.Group("/api/cms/", middlewares.JwtAuthMiddleware(utils.RoleUser, utils.RoleAdmin))

	// CMS User
	cmsRouteAdmin.GET("/users", userController.FindAll)
	cmsRouteAdmin.GET("/users/:id", userController.FindByID)
	cmsRouteAdmin.POST("/users", userController.Create)
	cmsRouteAdmin.PUT("/users/:id", userController.Update)
	cmsRouteAdmin.DELETE("/users/:id", userController.Delete)
	cmsRouteAllRole.PUT("/user/profile/:id", userController.UserUpdate)

	// CMS Role
	cmsRouteAdmin.GET("/roles", roleController.FindAll)
	cmsRouteAdmin.GET("/roles/:id", roleController.FindByID)
	cmsRouteAdmin.POST("/roles", roleController.Create)
	cmsRouteAdmin.PUT("/roles/:id", roleController.Update)
	cmsRouteAdmin.DELETE("/roles/:id", roleController.Delete)

	// CMS Order
	cmsRouteAllRole.GET("/orders", orderController.FindAll)
	cmsRouteAllRole.GET("/orders/:id", orderController.FindByID)
	cmsRouteAllRole.POST("/orders", orderController.Create)
	cmsRouteAllRole.PUT("/orders/:id", orderController.Update)
	cmsRouteAllRole.DELETE("/orders/:id", orderController.Delete)

	// CMS Transaction
	cmsRouteAllRole.GET("/transactions", transactionController.FindAll)
	cmsRouteAllRole.GET("/transactions/:id", transactionController.FindByID)
	cmsRouteAllRole.POST("/transactions", transactionController.Create)
	cmsRouteAllRole.PUT("/transactions/:id", transactionController.Update)
	cmsRouteAllRole.DELETE("/transactions/:id", transactionController.Delete)

	// CMS Invoice
	cmsRouteAllRole.GET("/invoices", transactionController.FindAll)
	cmsRouteAllRole.GET("/invoices/:id", transactionController.FindByID)
	cmsRouteAllRole.POST("/invoices", transactionController.Create)
	cmsRouteAllRole.PUT("/invoices/:id", transactionController.Update)
	cmsRouteAllRole.DELETE("/invoices/:id", transactionController.Delete)

	// Car
	cmsRouteAdmin.POST("/cars", carController.Create)
	r.GET("/api/cms/cars", carController.GetAll)
	r.GET("/api/cms/cars/:id", carController.GetByID)
	cmsRouteAdmin.GET("/cars/chart-data", carController.GetCarChartData)
	cmsRouteAdmin.PUT("/cars/:id", carController.Update)
	cmsRouteAdmin.DELETE("/cars/:id", carController.Delete)

	// BrandCar
	cmsRouteAdmin.POST("/brand-cars", brandCarController.Create)
	r.GET("/api/cms/brand-cars", brandCarController.GetAll)
	r.GET("/api/cms/brand-cars/:id", brandCarController.GetByID)
	cmsRouteAdmin.PUT("/brand-cars/:id", brandCarController.Update)
	cmsRouteAdmin.DELETE("/brand-cars/:id", brandCarController.Delete)

	// TypeCar
	cmsRouteAdmin.POST("/type-cars", typeCarController.Create)
	r.GET("/api/cms/type-cars", typeCarController.GetAll)
	r.GET("/api/cms/type-cars/:id", typeCarController.GetByID)
	cmsRouteAdmin.PUT("/type-cars/:id", typeCarController.Update)
	cmsRouteAdmin.DELETE("/type-cars/:id", typeCarController.Delete)

}
