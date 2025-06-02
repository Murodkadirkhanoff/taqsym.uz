package handler

import (
	"net/http"

	"github.com/Murodkadirkhanoff/taqsym.uz/internal/user/domain"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uc domain.Usecase
}

func NewUserHandler(uc domain.Usecase) *userHandler {
	return &userHandler{uc: uc}
}

func (h *userHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.GET("/", h.getAll)
}

func (h *userHandler) getAll(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.uc.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
