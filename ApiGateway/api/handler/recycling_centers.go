package handlers

import (
	pb "item_api/genproto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAddRecyclingCenter handles the creation of a new recycling center.
// @Summary Create Add Recycling Center
// @Description Create a new recycling center
// @Tags RecyclingCenter
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateAddRecyclingCenterRequest true "Create Add Recycling Center"
// @Success 200 {object} string "Recycling Center Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/recycling-center [post]
func (h *Handler) CreateAddRecyclingCenter(gn *gin.Context) {
	request := pb.CreateAddRecyclingCenterRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateAddRecyclingCenter request: %+v", &request)

	response, err := h.EcoService.CreateAddRecyclingCenter(gn, &request)
	if err != nil {
		log.Printf("Error creating recycling center: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Recycling center created successfully: %+v", response)
}

// SearchRecyclingCenter handles searching for recycling centers based on material.
// @Summary Search Recycling Center
// @Description Search for recycling centers based on material
// @Tags RecyclingCenter
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param material query string false "Material"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} string "Recycling Centers Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/recycling-centers [get]
func (h *Handler) SearchRecyclingCenter(gn *gin.Context) {
	material := gn.Query("material")

	page, err := strconv.Atoi(gn.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(gn.Query("limit"))
	if err != nil {
		limit = 10
	}

	request := pb.SearchRecyclingCenterRequest{
		Material: material,
		Page:     int32(page),
		Limit:    int32(limit),
	}
	log.Printf("Received SearchRecyclingCenter request: %+v", &request)

	response, err := h.EcoService.SearchRecyclingCenter(gn, &request)
	if err != nil {
		log.Printf("Error searching recycling centers: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Recycling centers retrieved successfully: %+v", response)
}
