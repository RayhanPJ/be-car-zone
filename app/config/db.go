package config

import (
	"be-car-zone/app/models"
	"be-car-zone/app/pkg/utils"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	dbProvider := utils.Getenv("DB_PROVIDER", "mysql")
	var db *gorm.DB

	if dbProvider == "postgres" {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")
		// production
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		log.Println("Database Connected at", dbProvider, "provider")

		db = dbGorm

	} else {
		username := utils.Getenv("DB_USERNAME", "root")
		password := utils.Getenv("DB_PASSWORD", "root")
		host := utils.Getenv("DB_HOST", "127.0.0.1")
		port := utils.Getenv("DB_PORT", "3306")
		database := utils.Getenv("DB_NAME", "db_name")

		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

		dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		log.Println("Database Connected at", dbProvider, "provider")

		db = dbGorm

	}

	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Invoice{},
		&models.Order{},
		&models.Transaction{},
		&models.Car{},
		&models.TypeCar{},
		&models.BrandCar{},
	)

	return db

}
