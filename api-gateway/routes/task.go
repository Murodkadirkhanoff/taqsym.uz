package routes

import (
	"net/http"

	"github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/grpc_clients"
	taskpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/task"
	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      int64  `json:"user_id" binding:"required"`
}

func TasksListHandler(c *gin.Context) {
	resp, err := grpc_clients.TaskClient.ListTasks(c, &taskpb.ListTasksRequest{})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверные учетные данные"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp.Tasks})
}

func CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := grpc_clients.TaskClient.Create(c, &taskpb.CreateTaskRequest{
		Title:       req.Title,
		Description: req.Description,
		UserId:      req.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать задачу"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": resp.Task})
}
