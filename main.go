package main

import (
	"log"
	"os"
	"time"

	"github.com/Niraj1910/school-census-go-backend/config"
	"github.com/Niraj1910/school-census-go-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load envoironment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	config.ConnectDB()

	// Connect cloudinary
	config.ConnectCloudinary()

	// Initialize the Gin router
	router := gin.Default()

	allowOrigins := []string{"http://localhost:3000"}
	if clientURL := os.Getenv("CLIENT_URL"); clientURL != "" {
		allowOrigins = append(allowOrigins, clientURL)
	}

	// CORS: middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-CSRF-Token", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Preflight result cache duration
	}))

	// School routes
	routes.RegisterSchoolRoutes(router)

	// Auth routes
	routes.RegisterAuthRoutes(router)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "School census API is running",
		})
	})

	// Start server
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s \n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
