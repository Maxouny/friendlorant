package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := gin.Default()

		corsConfig := cors.DefaultConfig()
		corsConfig.AllowHeaders = []string{"Accept"}
		corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

		r.Use(cors.New(corsConfig))
	}
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		authHeader := ctx.GetHeader("Authorization")
// 		if authHeader == "" {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Authorization required",
// 			})
// 			ctx.Abort()
// 			return
// 		}
// 		token := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
// 		if token == "" {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid token",
// 			})
// 			ctx.Abort()
// 			return
// 		}
// 		parsedToken, err := utils.ParseToken(token)

// 		if err != nil || !parsedToken.Valid {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid token",
// 			})
// 			ctx.Abort()
// 			return
// 		}
// 		claims, ok := parsedToken.Claims.(*utils.JWTClaims)
// 		if !ok || !parsedToken.Valid {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid token claims",
// 			})
// 			ctx.Abort()
// 			return
// 		}
// 		ctx.Set("id", claims.UserID)
// 		ctx.Next()
// 	}
// }
