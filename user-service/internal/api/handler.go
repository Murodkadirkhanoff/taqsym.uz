package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taqsym/taqsym.uz/user-service/internal/domain"
)

// Handler handles HTTP requests for the user service
type Handler struct {
	userUseCase domain.UserUseCase
}

// NewHandler creates a new Handler
func NewHandler(userUseCase domain.UserUseCase) *Handler {
	return &Handler{
		userUseCase: userUseCase,
	}
}

// RegisterRoutes registers the API routes
func (h *Handler) RegisterRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/health", h.HealthCheck)

	// API routes
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", h.GetAllUsers)
			users.GET("/:id", h.GetUserByID)
			users.POST("", h.CreateUser)
			users.PUT("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}

		// Authentication routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Login)
		}
	}
}

// HealthCheck handles health check requests
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetAllUsers handles requests to get all users
func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.userUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID handles requests to get a user by ID
func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser handles requests to create a new user
func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.userUseCase.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Don't return the password in the response
	user.Password = ""

	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles requests to update an existing user
func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.ID = id

	if err := h.userUseCase.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Don't return the password in the response
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// DeleteUser handles requests to delete a user
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.userUseCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

// Login handles user authentication
func (h *Handler) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.userUseCase.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Don't return the password in the response
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		// In a real implementation, we would generate and return a JWT token here
		"token": "sample-token",
	})
}
