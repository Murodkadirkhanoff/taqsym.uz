package routes

import (
	"fmt"
	"net/http"

	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/grpc_clients"
	authpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/auth"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := grpc_clients.AuthClient.Login(c, &authpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверные учетные данные"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp.Token})
}

func RegisterHandler(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := grpc_clients.AuthClient.Register(c, &authpb.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": resp.Id, "message": resp.Message})
}

func ProfileHandler(c *gin.Context) {
	userID := c.Value("userID").(int)
	fmt.Println(userID)
	resp, err := grpc_clients.AuthClient.Profile(c, &authpb.ProfileRequest{
		Id: int64(userID),
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": resp.Id, "name": resp.Name, "email": resp.Email})
}
