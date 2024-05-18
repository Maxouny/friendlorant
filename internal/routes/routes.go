package routes

import (
	"friendlorant/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userController *controllers.UserController) {
	// main route /api
	api := r.Group("/api")
	{
		// v1 route /api/v1
		v1 := api.Group("/v1")
		{
			setupUserRoutes(v1, userController)
		}
	}
}

// /api/v1/users
func setupUserRoutes(v1 *gin.RouterGroup, userController *controllers.UserController) {
	users := v1.Group("/users")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
		users.GET("/user/:id", userController.GetUserByID)
		users.GET("/user/email/:email", userController.GetUserByEmail)
		users.GET("/user/username/:username", userController.GetUserByUsername)
		users.PUT("/user/:id", userController.UpdateUser)
		users.DELETE("/user/:id", userController.DeleteUser)
	}
}
