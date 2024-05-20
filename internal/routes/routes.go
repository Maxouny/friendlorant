package routes

import (
	"friendlorant/internal/controllers"
	"friendlorant/internal/socket"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userController *controllers.UserController) {
	// main route /api
	// r.Use(middleware.AuthMiddleware())
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
		users.GET("/", userController.GetUsers)
		users.PUT("/user/:id", userController.UpdateUser)
		users.DELETE("/user/:id", userController.DeleteUser)
		// socket
		users.GET("/ws", func(ctx *gin.Context) { socket.HandleConnections(ctx.Writer, ctx.Request) })
	}
}
