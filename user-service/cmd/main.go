package main

import (
	"log"
	"net"

	authpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/auth"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/handler"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/repository"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/usecase"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/pkg/db"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	database, err := db.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	log.Println("Успешное подключение к базе данных!")

	repo := repository.NewUserRepo(database)
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)

	// r := router.SetupRouter(h)

	// port := os.Getenv("USER_SERVICE_APP_PORT")
	// port := "8081"
	// fmt.Println(port)
	// r.Run(":" + port)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("не удалось слушать: %v", err)
	}

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, h)
	reflection.Register(grpcServer)

	log.Println("AuthService запущен на порту :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("ошибка запуска gRPC сервера: %v", err)
	}
}
