package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taqsym/taqsym.uz/task-service/internal/domain"
)

// Handler handles HTTP requests for the task service
type Handler struct {
	taskUseCase domain.TaskUseCase
}

// NewHandler creates a new Handler
func NewHandler(taskUseCase domain.TaskUseCase) *Handler {
	return &Handler{
		taskUseCase: taskUseCase,
	}
}

// RegisterRoutes registers the API routes
func (h *Handler) RegisterRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/health", h.HealthCheck)

	// API routes
	api := router.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.GET("", h.GetAllTasks)
			tasks.GET("/:id", h.GetTaskByID)
			tasks.POST("", h.CreateTask)
			tasks.PUT("/:id", h.UpdateTask)
			tasks.DELETE("/:id", h.DeleteTask)
		}
	}
}

// HealthCheck handles health check requests
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetAllTasks handles requests to get all tasks
func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.taskUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID handles requests to get a task by ID
func (h *Handler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask handles requests to create a new task
func (h *Handler) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.taskUseCase.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// UpdateTask handles requests to update an existing task
func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.ID = id

	if err := h.taskUseCase.Update(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask handles requests to delete a task
func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := h.taskUseCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
