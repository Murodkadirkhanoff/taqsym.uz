package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/handler"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/repository"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/router"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to connect db:", err)
	}

	repo := repository.NewUserRepo(db)
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)

	r := router.SetupRouter(h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	r.Run(":" + port)
}
