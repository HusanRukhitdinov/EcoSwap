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
		users.POST("/reset_password", h.ResetPassword)
		users.POST("/refresh_token", h.RefreshToken)
		users.POST("/logout", h.Logout)
		users.GET("/users", h.GetProfile)
		users.PUT("/:id", h.EditProfile)
		users.GET("/", h.ListUsers)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("/eco_points", h.GetEcoPoints)
		users.POST("/eco_points", h.AddEcoPoints)
		users.GET("/eco_points_history", h.GetEcoPointsHistory)
	}
	return router
}
