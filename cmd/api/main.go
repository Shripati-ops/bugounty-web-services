package main

// This file is responsible for initializing the API server and how it interacts with the application.
import (
	"bugounty/internals/config"
	"bugounty/internals/db"
	"bugounty/internals/httpRoutes"
	"fmt"
)

func main() {
	config, err := config.LoadConfig()
	db, dbErr := db.NewPostgresDB(config.DB)
	go func() {
		httpRoutes.SetupRoutes(config.HTTP, db)
	}()

	if dbErr != nil {
		fmt.Printf("Failed to connect to database: %v", dbErr)
	}
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
	}
	fmt.Println("Database connected successfully", db)
	select {}
}

