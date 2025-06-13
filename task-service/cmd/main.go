package main

import (
	"log"

	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/pkg/db"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()
}
