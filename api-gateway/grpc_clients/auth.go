package grpc_clients

import (
	"log"

	pb "github.com/Murodkadirkhanoff/taqsym.uz/api-gateway/proto/generated/pb"

	"google.golang.org/grpc"
)

var AuthClient pb.AuthServiceClient

func InitAuthClient() {
	conn, err := grpc.Dial("auth-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("не удалось подключиться к auth-service: %v", err)
	}
	AuthClient = pb.NewAuthServiceClient(conn)
}
