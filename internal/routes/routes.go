package routes

import (
	"friendlorant/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userController *controllers.UserController) {
	v1 := r.Group("/api/v1")
	{

		// TODO: getAllusers,

		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		v1.GET("/user/:id", userController.GetUserByID)
		v1.GET("/user/email/:email", userController.GetUserByEmail) // Измененный путь
		v1.PUT("/user/:id", userController.UpdateUser)
		v1.DELETE("/user/:id", userController.DeleteUser)
	}
}
