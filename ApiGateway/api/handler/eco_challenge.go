package handlers

import (
	pb "item_api/genproto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEcoChallenge handles the creation of a new eco challenge.
// @Summary Create Eco Challenge
// @Description Create a new eco challenge
// @Tags EcoChallenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateEcoChallengeRequest true "Create Eco Challenge"
// @Success 200 {object} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/eco-challenges [post]
func (h *Handler) CreateEcoChallenge(gn *gin.Context) {
	request := pb.CreateEcoChallengeRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateEcoChallenge request: %+v", &request)

	response, err := h.EcoService.CreateEcoChallenge(gn, &request)
	if err != nil {
		log.Printf("Error creating eco challenge: %v", err)
		InternalServerError(gn, err)
		return
	}
	gn.JSON(http.StatusOK, response)
	log.Printf(`eco challenge created successfully: %+v`, response)
}
