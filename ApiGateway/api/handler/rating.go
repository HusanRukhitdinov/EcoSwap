package handlers

import (
	"fmt"
	pb "item_api/genproto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAddUserRating handles the creation of a new user rating.
// @Summary Create Add User Rating
// @Description Create a new user rating
// @Tags UserRating
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateAddUserRatingRequest true "Create Add User Rating"
// @Success 200 {object} string "User Rating Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user-ratings [post]
func (h *Handler) CreateAddUserRating(gn *gin.Context) {
	request := pb.CreateAddUserRatingRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateAddUserRating request: %+v", &request)

	response, err := h.EcoService.CreateAddUserRating(gn, &request)
	if err != nil {
		log.Printf("Error creating user rating: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("User rating created successfully: %+v", response)
}

// GetUserRating handles retrieving user ratings based on user ID.
// @Summary Get User Rating
// @Description Get user ratings based on user ID
// @Tags UserRating
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string true "User ID"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} string "User Ratings Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user-ratings [get]
func (h *Handler) GetUserRating(gn *gin.Context) {
	userID := gn.Query("user_id")
	if userID == "" {
		log.Printf("User ID is required")
		BadRequest(gn, fmt.Errorf("user_id is required"))
		return
	}

	page, err := strconv.Atoi(gn.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(gn.Query("limit"))
	if err != nil {
		limit = 10
	}

	request := pb.GetUserRatingRequest{
		UserId: userID,
		Page:   int32(page),
		Limit:  int32(limit),
	}
	log.Printf("Received GetUserRating request: %+v", &request)

	response, err := h.EcoService.GetUserRating(gn, &request)
	if err != nil {
		log.Printf("Error retrieving user ratings: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("User ratings retrieved successfully: %+v", response)
}
