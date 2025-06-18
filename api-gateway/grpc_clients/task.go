package grpc_clients

import (
	"log"

	taskpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/task"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var TaskClient taskpb.TaskServiceClient

func InitTaskClient() {
	conn, err := grpc.NewClient("task-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться к auth-service: %v", err)
	}
	TaskClient = taskpb.NewTaskServiceClient(conn)
}
