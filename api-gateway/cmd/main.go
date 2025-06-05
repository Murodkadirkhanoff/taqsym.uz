package main

import (
	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/grpc_clients"
	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	grpc_clients.InitAuthClient()

	r.POST("/login", routes.LoginHandler)

	r.Run(":8081")
}
