package api

import (
	handlers "item_api/api/handler"
	// middleware "item_api/api/middlewere"
	_ "item_api/docs"
	genproto "item_api/genproto"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// RouterApi @title API Service
// @version 1.0
// @description API service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func RouterApi(con1 *grpc.ClientConn, con2 *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	authCon := genproto.NewAuthServiceClient(con1)
	ecoCon := genproto.NewEcoServiceClient(con2)
	h := handlers.NewHandler(authCon, ecoCon)

	authRoutes := router.Group("/")
	// authRoutes.Use(middleware.MiddleWare())
	

		items := authRoutes.Group("/api/ecoSwap")
		{
			items.POST("/create", h.CreateItem)
			items.PUT("/:id", h.UpdateItem)
			items.DELETE("/:id", h.DeleteItem)
			items.GET("/", h.GetAllItems)
			items.GET("/:id", h.GetByIdItem)
			items.GET("/search", h.SearchItemsAndFilt)
			items.POST("/change_swap", h.CreateChangeSwap)
			items.PUT("/accept_swap/:id", h.UpdateAcceptSwap)
			items.PUT("/reject_swap/:id", h.UpdateRejactSwap)
			items.GET("/change_swap/:id", h.GetChangeSwap)
			items.POST("/recycling_center", h.CreateAddRecyclingCenter)
			items.GET("/recycling_center", h.SearchRecyclingCenter)
			items.POST("/recycling_submissions/:id", h.CreteRecyclingSubmissions)
			items.POST("/user_rating/:id", h.CreateAddUserRating)
			items.GET("/user_rating/:id", h.GetUserRating)
			items.POST("/item_category", h.CreateItemCategoryManag)
			items.GET("/statistics", h.GetStatistics)
			items.GET("/monitoring_user_activity/:id", h.GetMonitoringUserActivity)
			items.POST("/eco_challenge", h.CreateEcoChallenge)
			items.POST("/participate_challenge", h.CreateParticipateChallenge)
			items.PUT("/eco_challenge_result", h.UpdateEcoChallengeResult)
			items.POST("/eco_tips/:id", h.CreateAddEcoTips)
			items.GET("/eco_tips/:id", h.GetAddEcoTips)
		}
		return router
	}



