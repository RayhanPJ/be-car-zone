package api

import (
	"be-car-zone/app/config"
	// "be-car-zone/app/docs"
	"be-car-zone/app/routes"
	"be-car-zone/app/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	App *gin.Engine
	DB  *gorm.DB
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	App = gin.New()

	// Load environment variables
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file in main")
		}
	}

	// Set Swagger info
	// docs.SwaggerInfo.Title = "News API"
	// docs.SwaggerInfo.Description = "This is a sample server News."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:8080")
	// if environment == "development" {
	// 	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// } else {
	// 	docs.SwaggerInfo.Schemes = []string{"https"}
	// }

	// // Connect to the database
	DB = config.ConnectDataBase()

	// Setup routes
	routes.SetupRouter(DB, App)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
