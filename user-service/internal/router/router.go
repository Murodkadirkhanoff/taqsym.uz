package router

import (
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	return r
}
