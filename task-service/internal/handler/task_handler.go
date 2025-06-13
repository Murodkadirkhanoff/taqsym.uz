package handler

import (
	"context"
	"errors"

	taskpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/task"
	"github.com/Murodkadirkhanoff/taqsym.uz/task-service/internal/domain"
)

type TaskHandler struct {
	taskpb.UnimplementedTaskServiceServer
	uc domain.Usecase
}

func (h *TaskHandler) Create(ctx context.Context, request *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	task := domain.Task{
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		UserID:      request.GetUserId(),
	}
	h.uc.Create(ctx, &task)
	protoTask := &taskpb.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		UserId:      task.UserID,
	}

	return &taskpb.CreateTaskResponse{
		Task: protoTask,
	}, nil
}

func (h *TaskHandler) List(ctx context.Context, request *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.uc.List(ctx)

	protoTasks := []*taskpb.Task{}

	for _, task := range tasks {
		protoTask := &taskpb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			UserId:      task.UserID,
		}

		protoTasks = append(protoTasks, protoTask)
	}

	if err != nil {
		return nil, errors.New("can not fetch tasks list")
	}

	return &taskpb.ListTasksResponse{
		Tasks:   protoTasks,
		Message: "Tasks list fetched successfully",
	}, nil
}
