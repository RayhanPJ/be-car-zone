package main

import (
	"be-car-zone/app/api"
	"be-car-zone/app/pkg/utils"
	"log"
)

func main() {
	// Menentukan port untuk server
	port := utils.Getenv("HOST", "localhost:8080")

	// Menjalankan server HTTP
	log.Printf("Server running on http://localhost%s", port)
	err := api.App.Run(port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
