package main

import (
	"fmt"
	"log"
	"sistem-pembiayaan/config"
)

func main() {
	// Load config menggunakan viper
	viperConfig := config.InitConfig()

	// Init DB dengan config di env
	db, err := config.InitDB(viperConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Init Validator
	validate := config.NewValidator()

	// Init Gin
	app := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		Router:   app,
		Validate: validate,
		Viper:    viperConfig,
	})

	// Ambil port dari config
	port := viperConfig.GetString("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := app.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
