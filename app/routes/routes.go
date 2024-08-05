package routes

import (
	"be-car-zone/app/controllers"

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

	// Init controllers
	carController := &controllers.CarController{DB: db}
	brandCarController := &controllers.BrandCarController{DB: db}
	typeCarController := &controllers.TypeCarController{DB: db}

	// Car
	r.POST("/cars", carController.Create)
	r.GET("/cars", carController.GetAll)
	r.GET("/cars/:id", carController.GetByID)
	r.PUT("/cars/:id", carController.Update)
	r.DELETE("/cars/:id", carController.Delete)

	// BrandCar
	r.POST("/brand-cars", brandCarController.Create)
	r.GET("/brand-cars", brandCarController.GetAll)
	r.GET("/brand-cars/:id", brandCarController.GetByID)
	r.PUT("/brand-cars/:id", brandCarController.Update)
	r.DELETE("/brand-cars/:id", brandCarController.Delete)

	// TypeCar
	r.POST("/type-cars", typeCarController.Create)
	r.GET("/type-cars", typeCarController.GetAll)
	r.GET("/type-cars/:id", typeCarController.GetByID)
	r.PUT("/type-cars/:id", typeCarController.Update)
	r.DELETE("/type-cars/:id", typeCarController.Delete)

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	authController := &controllers.AuthController{DB: db}

	// Authentication User
	authRoute := r.Group("/api/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)

}