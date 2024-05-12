package router

import (
	"friendlorant/internal/controllers"
	"friendlorant/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userRepo repositories.UserRepository) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		userController := controllers.NewUserController(userRepo)

		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		v1.GET("/user/:id", userController.GetUserByID)
		v1.GET("/user/:email", userController.GetUserByEmail)
		v1.PUT("/user/:id", userController.UpdateUser)
		v1.DELETE("/user/:id", userController.DeleteUser)
	}
	return r
}
