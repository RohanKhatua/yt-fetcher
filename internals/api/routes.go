package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rohankhatua/yt-fetcher/internals/api/handlers"
	"github.com/rohankhatua/yt-fetcher/internals/services"
)

func SetupRoutes(r *gin.Engine, userService *services.UserService) {
	userHandler := handlers.NewUserHandler(userService)

	r.Use(gin.Logger())

	r.POST("/users", userHandler.Create)
	r.GET("/users/:id", userHandler.GetByID)
	r.PUT("/users", userHandler.Update)
	r.DELETE("/users/:id", userHandler.Delete)
}