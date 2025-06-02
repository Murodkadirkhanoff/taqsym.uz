package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Murodkadirkhanoff/taqsym.uz/internal/user/handler"
	"github.com/Murodkadirkhanoff/taqsym.uz/internal/user/repository"
	"github.com/Murodkadirkhanoff/taqsym.uz/internal/user/usecase"
	"github.com/Murodkadirkhanoff/taqsym.uz/pkg/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка конфигурации
	dsn, exists := os.LookupEnv("DATABASE_DSN")
	if !exists {
		log.Fatal("DATABASE_DSN not set in environment")
	}

	// Подключение к БД
	pg, err := db.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	defer pg.Close()

	// Инициализация слоёв
	userRepo := repository.NewUserRepository(pg)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Настройка роутера
	r := gin.Default()
	userHandler.RegisterRoutes(r)

	// Graceful shutdown
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
