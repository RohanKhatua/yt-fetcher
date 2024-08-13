package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohankhatua/yt-fetcher/internals/api"
	"github.com/rohankhatua/yt-fetcher/internals/config"
	"github.com/rohankhatua/yt-fetcher/internals/db"
	"github.com/rohankhatua/yt-fetcher/internals/services"
)

func main () {
	// Load Configuration for database
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database with configuration
	db.Init(cfg)

	// Initialize user service which is used by the user handler
	userService := &services.UserService{}

	// Initialize the router and setup routes
	router := gin.Default()
	api.SetupRoutes(router, userService)

	// Run the server
	log.Printf("Starting server on port %d", cfg.ApplicationPort)

	if err := router.Run(":" + strconv.Itoa(cfg.ApplicationPort)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}