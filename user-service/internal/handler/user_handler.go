package handler

import (
	"context"
	"errors"

	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/domain"

	pb "github.com/Murodkadirkhanoff/taqsym.uz/user-service/proto/generated/pb"
)

type UserHandler struct {
	pb.UnimplementedAuthServiceServer
	uc domain.UserUseCase
}

// type UserHandler struct {
// 	uc domain.UserUseCase
// }

func NewUserHandler(uc domain.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	// return nil, errors.New(err.Error())
	// 	return nil
	// }

	token, err := h.uc.Login(ctx, input.Email, input.Password)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &pb.LoginResponse{Token: token}, nil
}

// func (h *UserHandler) Register(c *gin.Context) {
// 	var user domain.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.uc.Register(c.Request.Context(), &user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
// }

// func (h *UserHandler) Login(c *gin.Context) {
// 	var input struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	token, err := h.uc.Login(c.Request.Context(), input.Email, input.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

// func (h *UserHandler) Profile(c *gin.Context) {
// 	userID := c.Value("userID").(int)

// 	user, err := h.uc.Profile(c, userID)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// TODO change response to DTO
// 	c.JSON(http.StatusOK, gin.H{
// 		"user": gin.H{
// 			"id":    user.ID,
// 			"name":  user.Name,
// 			"email": user.Email,
// 		},
// 	})
// }
