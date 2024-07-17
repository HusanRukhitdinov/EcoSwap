package handlers

import (
	"fmt"
	pb "item_api/genproto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func IsValidLimit(limit string) (int, error) {
	if limit == "" {
		return 0, nil
	}
	return strconv.Atoi(limit)
}

func IsValidOffset(offset string) (int, error) {
	if offset == "" {
		return 0, nil
	}
	return strconv.Atoi(offset)
}

// CreateItem handles the creation of a new item.
// @Summary Create Item
// @Description Create a new item
// @Tags Item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.CreateItemRequest true "Create Item"
// @Success 200 {object} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/item [post]
func (h *Handler) CreateItem(gn *gin.Context) {
	request := pb.CreateItemRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received CreateItem request: %+v", &request)

	response, err := h.EcoService.CreateItem(gn, &request)
	if err != nil {
		log.Printf("Error creating item: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Item created successfully: %+v", response)
}

// UpdateItem handles updating an existing item.
// @Summary Update Item
// @Description Update an existing item
// @Tags Item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Item ID"
// @Param Create body genproto.UpdateItemRequest true "Update Item"
// @Success 200 {object} string "Item Updated"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/item/{id} [put]
func (h *Handler) UpdateItem(gn *gin.Context) {
	itemID := gn.Param("id")
	if !IsValidUUID(itemID) {
		log.Printf("Invalid UUID: %s", itemID)
		BadRequest(gn, fmt.Errorf("id is not valid"))
		return
	}
	
	request := pb.UpdateItemRequest{Id: itemID}
	if err := gn.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(gn, err)
		return
	}
	log.Printf("Received UpdateItem request: %+v", &request)

	response, err := h.EcoService.UpdateItem(gn, &request)
	if err != nil {
		log.Printf("Error updating item: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Item updated successfully: %+v", response)
}

// DeleteItem handles deleting an item.
// @Summary Delete Item
// @Description Delete an item
// @Tags Item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param item_id path string true "Item ID"
// @Success 200 {string} string "Item Deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/item/{item_id} [delete]
func (h *Handler) DeleteItem(gn *gin.Context) {
	itemID := gn.Param("item_id")
	if !IsValidUUID(itemID) {
		log.Printf("Invalid UUID: %s", itemID)
		BadRequest(gn, fmt.Errorf("item_id is not valid"))
		return
	}

	request := pb.DeleteItemRequest{ItemId: itemID}

	log.Printf("Received DeleteItem request for item ID: %s", itemID)

	response, err := h.EcoService.DeleteItem(gn, &request)
	if err != nil {
		log.Printf("Error deleting item: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Item deleted successfully: %+v", response)
}

// GetAllItems handles retrieving all items.
// @Summary Get All Items
// @Description Get all items
// @Tags Item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "Item ID"
// @Param name query string false "Item Name"
// @Param category_id query string false "Category ID"
// @Param condition query string false "Condition"
// @Param owner_id query string false "Owner ID"
// @Param status query string false "Status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} string "Items Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/items [get]
func (h *Handler) GetAllItems(gn *gin.Context) {
	limit, err := IsValidLimit(gn.Query("limit"))
	if err != nil {
		log.Printf("Invalid limit: %v", err)
		BadRequest(gn, err)
		return
	}

	offset, err := IsValidOffset(gn.Query("offset"))
	if err != nil {
		log.Printf("Invalid offset: %v", err)
		BadRequest(gn, err)
		return
	}

	request := pb.GetAllItemsRequest{
		Id:         gn.Query("id"),
		Name:       gn.Query("name"),
		CategoryId: gn.Query("category_id"),
		Condition:  gn.Query("condition"),
		OwnerId:    gn.Query("owner_id"),
		Status:     gn.Query("status"),
		Limit:      int64(limit),
		Offset:     int64(offset),
	}
	log.Printf("Received GetAllItems request: %+v", &request)

	response, err := h.EcoService.GetAllItems(gn, &request)
	if err != nil {
		log.Printf("Error retrieving items: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Items retrieved successfully: %+v", response)
}

// GetByIdItem handles retrieving an item by ID.
// @Summary Get Item By ID
// @Description Get an item by ID
// @Tags Item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param item_id path string true "Item ID"
// @Success 200 {object} string "Item Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/item/{item_id} [get]
func (h *Handler) GetByIdItem(gn *gin.Context) {
	itemID := gn.Param("item_id")
	if !IsValidUUID(itemID) {
		log.Printf("Invalid UUID: %s", itemID)
		BadRequest(gn, fmt.Errorf("item_id is not valid"))
		return
	}

	request := pb.GetByIdItemRequest{ItemId: itemID}
	log.Printf("Received GetByIdItem request for item ID: %s", itemID)

	response, err := h.EcoService.GetByIdItem(gn, &request)
	if err != nil {
		log.Printf("Error retrieving item: %v", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
	log.Printf("Item retrieved successfully: %+v", response)
}

// SearchItemsAndFilt handles searching and filtering items.
// @Summary Search and Filter Items
// @Description Search and filter items
// @Tags Item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query string false "Search Query"
// @Param category query string false "Category"
// @Param condition query string false "Condition"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} string "Items Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/items/search [get]
func (h *Handler) SearchItemsAndFilt(gn *gin.Context) {
	page, err := IsValidLimit(gn.Query("page"))
	if err != nil {
		log.Printf("Invalid page: %v", err)
		BadRequest(gn, err)
		return
	}

	limit, err := IsValidLimit(gn.Query("limit"))
	if err != nil {
		log.Printf("Invalid limit: %v", err)
		BadRequest(gn, err)
		return
	}

	request := pb.SearchItemsAndFiltRequest{
		Query:     gn.Query("query"),
		Category:  gn.Query("category"),
		Condition: gn.Query("condition"),
		Page:      int32(page),
		Limit:     int32(limit),
	}
	log.Printf("Received SearchItemsAndFilt request: %+v", &request)

	response, err := h.EcoService.SearchItemsAndFilt(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}
