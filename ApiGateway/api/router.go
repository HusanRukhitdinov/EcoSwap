package api

import (
	handlers "item_api/api/handler"
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
	

		items := authRoutes.Group("/api")
		{
			items.POST("/create", h.CreateItem)
			items.PUT("/update:id", h.UpdateItem)
			items.DELETE("/delete:id", h.DeleteItem)
			items.GET("/getall", h.GetAllItems)
			items.GET("/getid:id", h.GetByIdItem)
			items.GET("/search", h.SearchItemsAndFilt)
			items.POST("/change-swap", h.CreateChangeSwap)
			items.PUT("/accept-swap/:id", h.UpdateAcceptSwap)
			items.PUT("/reject-swap/:id", h.UpdateRejactSwap)
			items.GET("/change-swap/:id", h.GetChangeSwap)
			items.POST("/recycling-center", h.CreateAddRecyclingCenter)
			items.GET("/recycling-center", h.SearchRecyclingCenter)
			items.POST("/recycling-submissions/:id", h.CreteRecyclingSubmissions)
			items.POST("/user-rating/:id", h.CreateAddUserRating)
			items.GET("/user-rating/:id", h.GetUserRating)
			items.POST("/item-category", h.CreateItemCategoryManag)
			items.GET("/statistics", h.GetStatistics)
			items.GET("/monitoring-user-activity/:id", h.GetMonitoringUserActivity)
			items.POST("/eco-challenge", h.CreateEcoChallenge)
			items.POST("/participate-challenge", h.CreateParticipateChallenge)
			items.PUT("/eco-challenge-result", h.UpdateEcoChallengeResult)
			items.POST("/eco-tips/:id", h.CreateAddEcoTips)
			items.GET("/eco-tips/:id", h.GetAddEcoTips)
		}
		return router
	}



