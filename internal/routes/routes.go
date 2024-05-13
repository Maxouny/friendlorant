package router

import (
	"friendlorant/internal/controllers"
	"friendlorant/internal/middleware"
	"friendlorant/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userRepo repositories.UserRepository) {
	r.Use(middleware.AuthMiddleware())

	v1 := r.Group("/api/v1")
	{
		userController := controllers.NewUserController(userRepo)

		// TODO: getAllusers,

		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		v1.GET("/user/:id", userController.GetUserByID)
		v1.GET("/user/email/:email", userController.GetUserByEmail) // Измененный путь
		v1.PUT("/user/:id", userController.UpdateUser)
		v1.DELETE("/user/:id", userController.DeleteUser)
	}
}
