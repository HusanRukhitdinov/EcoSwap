package handlers

import (
	pb "item_api/genproto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateItemCategoryManag handles the creation of a new item category.
// @Summary Create Item Category
// @Description Create a new item category
// @Tags ItemCategory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateItemCategoryManagRequest true "Create Item Category"
// @Success 200 {object} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/item-categories [post]
func (h *Handler) CreateItemCategoryManag(gn *gin.Context) {
	request := pb.CreateItemCategoryManagRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateItemCategoryManag request: %+v", &request)

	response, err := h.EcoService.CreateItemCategoryManag(gn, &request)
	if err != nil {
		log.Printf("Error creating item category: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Item category created successfully: %+v", response)
}

// GetStatistics handles retrieving statistics based on date range.
// @Summary Get Statistics
// @Description Get statistics based on date range
// @Tags Statistics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param startDate query string true "Start Date"
// @Param endDate query string true "End Date"
// @Success 200 {object} string "Statistics Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/statistics [get]
func (h *Handler) GetStatistics(gn *gin.Context) {
	startDate := gn.Query("startDate")
	endDate := gn.Query("endDate")

	request := pb.GetStatisticsRequest{
		StartDate: startDate,
		EndDate:   endDate,
	}
	log.Printf("Received GetStatistics request: %+v", &request)

	response, err := h.EcoService.GetStatistics(gn, &request)
	if err != nil {
		log.Printf("Error retrieving statistics: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Statistics retrieved successfully: %+v", response)
}

// GetMonitoringUserActivity handles retrieving user activity based on user ID and date range.
// @Summary Get Monitoring User Activity
// @Description Get user activity based on user ID and date range
// @Tags UserActivity
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param userId query string true "User ID"
// @Param startDate query string true "Start Date"
// @Param endDate query string true "End Date"
// @Success 200 {object} string "User Activity Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user-activity [get]
func (h *Handler) GetMonitoringUserActivity(gn *gin.Context) {
	userID := gn.Query("userId")
	startDate := gn.Query("startDate")
	endDate := gn.Query("endDate")

	request := pb.GetMonitoringUserActivityRequest{
		UserId:    userID,
		StartDate: startDate,
		EndDate:   endDate,
	}
	log.Printf("Received GetMonitoringUserActivity request: %+v", &request)

	response, err := h.EcoService.GetMonitoringUserActivity(gn, &request)
	if err != nil {
		log.Printf("Error retrieving user activity: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("User activity retrieved successfully: %+v", response)
}
