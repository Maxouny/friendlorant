package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"friendlorant/internal/config"
	"friendlorant/internal/controllers"
	"friendlorant/internal/database"
	"friendlorant/internal/repositories"
	"friendlorant/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	envCfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	dbConn, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer dbConn.Close(context.Background())

	userRepo := repositories.NewUserRepository(dbConn)

	userController := controllers.NewUserController(userRepo)

	r := gin.Default()

	routes.SetupRouter(r, userController)

	server := &http.Server{
		Addr:    ":" + envCfg.Port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server: ", err)
		}
	}()
	log.Println("Server started on port " + envCfg.Port)

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
