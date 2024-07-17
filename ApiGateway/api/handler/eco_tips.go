package handlers

import (
	pb "item_api/genproto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAddEcoTips handles the creation of a new eco tip.
// @Summary Create Eco Tip
// @Description Create a new eco tip
// @Tags EcoTips
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateAddEcoTipsRequest true "Create Eco Tip"
// @Success 200 {object} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/eco-tips [post]
func (h *Handler) CreateAddEcoTips(gn *gin.Context) {
	request := pb.CreateAddEcoTipsRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateAddEcoTips request: %+v", &request)

	response, err := h.EcoService.CreateAddEcoTips(gn, &request)
	if err != nil {
		log.Printf("Error creating eco tip: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Eco tip created successfully: %+v", response)
}

// GetAddEcoTips handles retrieving eco tips with pagination.
// @Summary Get Eco Tips
// @Description Get eco tips with pagination
// @Tags EcoTips
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Get query genproto.GetAddEcoTipsRequest true "Get Eco Tips"
// @Success 200 {object} string "Eco Tips Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/eco-tips [get]
func (h *Handler) GetAddEcoTips(gn *gin.Context) {
	request := pb.GetAddEcoTipsRequest{}
	if err := gn.ShouldBindQuery(&request); err != nil {
		log.Printf("Error binding query parameters: %v", err)
		BadRequest(gn, err)
		return
	}
	
	log.Printf("Received GetAddEcoTips request: %+v", &request)

	response, err := h.EcoService.GetAddEcoTips(gn, &request)
	if err != nil {
		log.Printf("Error retrieving eco tips: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Eco tips retrieved successfully: %+v", response)
}
