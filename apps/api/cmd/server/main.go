package main

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/fanik05/tcg-pocket-api/internal/models"
	"github.com/fanik05/tcg-pocket-api/internal/engine"
)

func main() {
	r := gin.Default()

	// 1. Setup CORS
	// This allows your Next.js app (on localhost:3000) to talk to this API
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 2. Health Check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "TCG Pocket API is live"})
	})

	// 3. Pairing Endpoint
	r.POST("/pair", func(c *gin.Context) {
		var players []models.Player

		// Bind JSON body to the slice
		if err := c.ShouldBindJSON(&players); err != nil {
			c.JSON(400, gin.H{"error": "Invalid player data provided"})
			return
		}

		// Use the logic you tested earlier
		pairing.SortPlayersBySwiss(players)
		matches := pairing.GenerateMatches(players)

		c.JSON(200, matches)
	})

	// 4. Start Server
	r.Run(":8080")
}