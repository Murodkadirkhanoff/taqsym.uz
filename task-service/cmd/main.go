package main

import (
	"log"
	"net"

	taskpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/task"
	"github.com/Murodkadirkhanoff/taqsym.uz/task-service/internal/handler"
	"github.com/Murodkadirkhanoff/taqsym.uz/task-service/internal/repository"
	"github.com/Murodkadirkhanoff/taqsym.uz/task-service/internal/usecase"
	"github.com/Murodkadirkhanoff/taqsym.uz/task-service/pkg/db"
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

	repo := repository.NewTaskRepository(database)
	uc := usecase.NewTaskUsecase(repo)
	h := handler.NewTaskHandler(uc)

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
	taskpb.RegisterTaskServiceServer(grpcServer, h)
	reflection.Register(grpcServer)

	log.Println("AuthService запущен на порту :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("ошибка запуска gRPC сервера: %v", err)
	}
}
