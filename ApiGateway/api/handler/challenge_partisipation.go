package handlers

import (
	pb "item_api/genproto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateParticipateChallenge handles the participation in an eco challenge.
// @Summary Participate in Eco Challenge
// @Description Participate in an eco challenge
// @Tags EcoChallenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateParticipateChallengeRequest true "Participate in Eco Challenge"
// @Success 200 {object} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/eco-challenges/participate [post]
func (h *Handler) CreateParticipateChallenge(gn *gin.Context) {
	request := pb.CreateParticipateChallengeRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateParticipateChallenge request: %+v", &request)
	response, err := h.EcoService.CreateParticipateChallenge(gn, &request)
	if err != nil {
		log.Printf("Error creating participate challenge: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Participate challenge created successfully: %+v", response)
}
// UpdateEcoChallengeResult handles updating the result of an eco challenge.
// @Summary Update Eco Challenge Result
// @Description Update the result of an eco challenge
// @Tags EcoChallenge
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body genproto.UpdateEcoChallengeRresultRequest true "Update Eco Challenge Result"
// @Success 200 {object} string "Result Updated"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/eco-challenges/result [put]
func (h *Handler) UpdateEcoChallengeResult(gn *gin.Context) {
	request := pb.UpdateEcoChallengeRresultRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received UpdateEcoChallengeResult request: %+v", &request)

	response, err := h.EcoService.UpdateEcoChallengeResult(gn, &request)
	if err != nil {
		log.Printf("Error updating eco challenge result: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Eco challenge result updated successfully: %+v", response)
}
