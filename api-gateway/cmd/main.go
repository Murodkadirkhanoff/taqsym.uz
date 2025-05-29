package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration
	if err := loadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Register routes
	registerRoutes(router)

	// Create HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("server.port")),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting API Gateway server on port %s", viper.GetString("server.port"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}

func registerRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// User service routes
		userAPI := api.Group("/users")
		{
			userAPI.GET("", proxyToService("user-service", "/api/users"))
			userAPI.GET("/:id", proxyToService("user-service", "/api/users/:id"))
			userAPI.POST("", proxyToService("user-service", "/api/users"))
			userAPI.PUT("/:id", proxyToService("user-service", "/api/users/:id"))
			userAPI.DELETE("/:id", proxyToService("user-service", "/api/users/:id"))
		}

		// Task service routes
		taskAPI := api.Group("/tasks")
		{
			taskAPI.GET("", proxyToService("task-service", "/api/tasks"))
			taskAPI.GET("/:id", proxyToService("task-service", "/api/tasks/:id"))
			taskAPI.POST("", proxyToService("task-service", "/api/tasks"))
			taskAPI.PUT("/:id", proxyToService("task-service", "/api/tasks/:id"))
			taskAPI.DELETE("/:id", proxyToService("task-service", "/api/tasks/:id"))
		}
	}
}

// proxyToService returns a handler that forwards requests to the specified service
func proxyToService(serviceName, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// In a real implementation, this would use a proper HTTP client to forward the request
		// For now, we'll just return a placeholder response
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Request proxied to %s%s", serviceName, path),
		})
	}
}
