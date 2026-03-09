package main

import (
	"log"

	"github.com/Derikklok/go-college-server/internal/config"
	"github.com/Derikklok/go-college-server/internal/database"
	"github.com/Derikklok/go-college-server/internal/student"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	database.Connect()

	student.Migrate()

	router := gin.Default()

	api := router.Group("/api")

	student.RegisterRoutes(api)

	log.Println("Server running on http://localhost:", config.AppPort)

	router.Run(":" + config.AppPort)
}
