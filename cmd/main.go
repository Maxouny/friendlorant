package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"friendlorant/db"
	"friendlorant/internal/config"
	"friendlorant/internal/repositories"
	routes "friendlorant/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer dbConn.Close()

	if err := db.Migrate(dbConn); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	userRepo := repositories.NewUserRepository(dbConn)

	r := gin.Default()

	routes.SetupRouter(r, userRepo)

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server: ", err)
		}
	}()
	log.Println("Server started on port " + cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shutdown server: ", err)
	}
	log.Println("Server stopped")
}
