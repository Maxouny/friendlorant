package main

import (
	"log"
	"net/http"

	"friendlorant/db"
	"friendlorant/internal/config"
	"friendlorant/internal/repositories"
	router "friendlorant/internal/routes"
)

func main() {
	cfg := config.LoadConfig()

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	r := router.SetupRouter(userRepo)

	log.Printf("Listening on port %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, r))
}
