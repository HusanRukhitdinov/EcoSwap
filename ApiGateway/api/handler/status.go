package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OK(gn *gin.Context) {
	gn.Header("Content-Type", "application/json")
	gn.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"time":    time.Now(),
		"success": true,
	})
}

func Created(gn *gin.Context) {
	gn.Header("Content-Type", "application/json")
	gn.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"time":    time.Now(),
		"success": true,
	})
}

func InternalServerError(gn *gin.Context, err error) {
	gn.Header("Content-Type", "application/json")
	gn.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"time":    time.Now(),
		"message": err.Error(),
		"success": false,
	})
	fmt.Println("Internal server error:", err)
}

func BadRequest(gn *gin.Context, err error) {
	gn.Header("Content-Type", "application/json")
	gn.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"time":    time.Now(),
		"message": err.Error(),
		"success": false,
	})
}

func Parse(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func IsLimitOffsetValidate(limit string) (int, error) {
	if limit == "" {
		return 0, nil
	}
	return strconv.Atoi(limit)
}

func IsAmount(id string) bool {
	_, err := strconv.ParseFloat(id, 32)
	return err == nil
}
