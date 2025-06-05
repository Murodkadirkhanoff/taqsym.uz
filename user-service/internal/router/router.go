package router

import (
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.UserHandler) *gin.Engine {
	r := gin.Default()
	// authRoutes := r.Group("/", middleware.AuthMiddleware())

	// r.POST("/register", h.Register)
	// r.POST("/login", h.Login)

	// authRoutes.GET("/profile", h.Profile)
	return r
}
