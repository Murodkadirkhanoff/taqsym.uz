package main

import (
	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/grpc_clients"
	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/middleware"
	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authRoutes := r.Group("/", middleware.AuthMiddleware())

	grpc_clients.InitAuthClient()
	grpc_clients.InitTaskClient()

	r.POST("/login", routes.LoginHandler)
	r.POST("/register", routes.RegisterHandler)
	r.GET("/tasks", routes.TasksListHandler)
	r.POST("tasks", routes.CreateTask)

	authRoutes.GET("/profile", routes.ProfileHandler)

	r.Run(":8081")
}
