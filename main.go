package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"Executor/routes"
	"Executor/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
		} else {
			log.Println("Successfully loaded .env file")
		}
	}
	go func() {
		fmt.Println("Starting Kafka consumer...")
		services.DataConsumer()
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port
	}

	e1 := gin.New()
	configCors := cors.DefaultConfig()
	//Apply CORS middleware
	// // Update with your allowed origins
	configCors.AllowAllOrigins = true
	configCors.AllowCredentials = true
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	configCors.AllowHeaders = []string{"*"}
	e1.Use(gzip.Gzip(gzip.DefaultCompression))
	e1.Use(cors.New(configCors))
	e1.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gateway!"})
	})
	e1.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gateway!"})
	})
	e1.Use(func(c *gin.Context) {
		if c.Request.TLS != nil {
			c.Writer.Header().Set("Strict-Transport-Security", fmt.Sprintf("max-age=%d; includeSubDomains; preload", int((60*24*time.Hour).Seconds())))
		}
		c.Next()
	})
	e1.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Security-Policy", "script-src 'self' 'unsafe-inline' https://cdn.jsdelivr.net/npm/swagger-ui-dist@<version>/swagger-ui-bundle.js'; img-src 'self'")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Next()
	})
	routes.LoadExecutorRequestRoute(e1)
	e1.Run(":" + port)
}
