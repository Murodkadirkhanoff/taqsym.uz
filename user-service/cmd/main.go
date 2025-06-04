package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/handler"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/repository"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/router"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/usecase"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/pkg/db"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе: %v", err)
	}
	defer database.Close()
	log.Println("Успешное подключение к базе данных!")

	repo := repository.NewUserRepo(database)
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)

	r := router.SetupRouter(h)

	port := os.Getenv("USER_SERVICE_APP_PORT")
	fmt.Println(port)
	r.Run(":" + port)
}
