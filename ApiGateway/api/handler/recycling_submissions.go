package handlers

import (
	pb "item_api/genproto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreteRecyclingSubmissions handles the creation of a new recycling submission.
// @Summary Create Recycling Submission
// @Description Create a new recycling submission
// @Tags RecyclingSubmission
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreteRecyclingSubmissionsRequest true "Create Recycling Submission"
// @Success 200 {object} string "Recycling Submission Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/recycling-submission [post]
func (h *Handler) CreteRecyclingSubmissions(gn *gin.Context) {
	request := pb.CreteRecyclingSubmissionsRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v",err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateRecyclingSubmissions request: %+v", &request)

	response, err := h.EcoService.CreateRecyclingSubmissions(gn, &request)
	if err != nil {
		log.Printf("Error creating recycling submission: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Recycling submission created successfully: %+v", response)
}
