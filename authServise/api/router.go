package api

import (
	handlers "eco_system/api/handler"
	_ "eco_system/docs"
	"eco_system/service"
"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterApi @title API Service
// @version 1.0
// @description API service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func RouterApi(authService *service.UserService) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	h := handlers.NewHandler(authService)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	users := router.Group("/api/user")
	{
		users.POST("/register", h.RegisterUser)
		users.POST("/login", h.LoginUser)
		users.POST("/reset-password", h.ResetPassword)
		users.POST("/refresh-token", h.RefreshToken)
		users.POST("/logout", h.Logout)
		users.GET("/profile/:user_id", h.GetProfile)
		users.PUT("/profile/:user_id", h.EditProfile)
		users.GET("/users", h.ListUsers)
		users.DELETE("/:user_id", h.DeleteUser)
		users.GET("/:user_id/eco-points", h.GetEcoPoints)
		users.POST("/eco-points", h.AddEcoPoint)
		users.GET("/eco-points/history:user_id", h.GetEcoPointsHistory)
	}
	return router
}
