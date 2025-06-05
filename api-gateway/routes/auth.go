package routes

import (
	"net/http"

	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/grpc_clients"
	pb "github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/proto/generated/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := grpc_clients.AuthClient.Login(c, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверные учетные данные"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp.Token})
}
