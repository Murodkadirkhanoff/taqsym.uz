package grpc_clients

import (
	"log"

	authpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthClient authpb.AuthServiceClient

func InitAuthClient() {
	conn, err := grpc.NewClient("user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться к auth-service: %v", err)
	}
	AuthClient = authpb.NewAuthServiceClient(conn)
}
