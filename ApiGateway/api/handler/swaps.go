package handlers

import (
	"fmt"
	pb "item_api/genproto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateChangeSwap handles the creation of a new swap request.
// @Summary Create Change Swap
// @Description Create a new swap request
// @Tags Swap
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateChangeSwapRequest true "Create Change Swap"
// @Success 200 {object} string "Swap Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/swap [post]
func (h *Handler) CreateChangeSwap(gn *gin.Context) {
	request := pb.CreateChangeSwapRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	response, err := h.EcoService.CreateChangeSwap(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// UpdateAcceptSwap handles accepting a swap request.
// @Summary Accept Swap
// @Description Accept a swap request
// @Tags Swap
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param swap_id path string true "Swap ID"
// @Success 200 {object} string "Swap Accepted"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/swap/{swap_id}/accept [put]
func (h *Handler) UpdateAcceptSwap(gn *gin.Context) {
	swapID := gn.Param("swap_id")
	if swapID == "" {
		BadRequest(gn, fmt.Errorf("swap_id is required"))
		return
	}

	request := pb.UpdateAcceptSwapRequest{SwapId: swapID}
	response, err := h.EcoService.UpdateAcceptSwap(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// UpdateRejactSwap handles rejecting a swap request.
// @Summary Reject Swap
// @Description Reject a swap request
// @Tags Swap
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param swap_id path string true "Swap ID"
// @Param Create body genproto.UpdateRejactSwapRequest true "Reject Swap"
// @Success 200 {object} string "Swap Rejected"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/swap/{swap_id}/reject [put]
func (h *Handler) UpdateRejactSwap(gn *gin.Context) {
	swapID := gn.Param("swap_id")
	if swapID == "" {
		BadRequest(gn, fmt.Errorf("swap_id is required"))
		return
	}

	request := pb.UpdateRejactSwapRequest{SwapId: swapID}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	response, err := h.EcoService.UpdateRejectSwap(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// GetChangeSwap handles retrieving swap requests based on status.
// @Summary Get Change Swap
// @Description Get swap requests based on status
// @Tags Swap
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param status query string false "Status"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} string "Swaps Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/swaps [get]
func (h *Handler) GetChangeSwap(gn *gin.Context) {
	status := gn.Query("status")

	page, err := strconv.Atoi(gn.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(gn.Query("limit"))
	if err != nil {
		limit = 10
	}

	request := pb.GetChangeSwapRequest{
		Status: status,
		Page:   int32(page),
		Limit:  int32(limit),
	}

	response, err := h.EcoService.GetChangeSwap(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}
