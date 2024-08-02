package main

import (
	"log"

	"be-car-zone/app/config"
	"be-car-zone/app/pkg/utils"
	"be-car-zone/app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title API CarZone E-Commerce
// @version 1.0
// @description This is a sample server for carzone e-commerce API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	var app *gin.Engine

	app = gin.New()

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db := config.ConnectDataBase()

	routes.SetupRouter(db, app)

	app.Run() // listen and serve on http://localhost:8080
}
