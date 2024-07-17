package handlers

import (
	pb "eco_system/genproto"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

// RegisterUser handles the creation of a new user.
// @Summary Register User
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Create body genproto.RegisterRequest true "Register"
// @Success 200 {string} string "Register Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/register [post]
func (h *Handler) RegisterUser(gn *gin.Context) {
	request := pb.RegisterRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		fmt.Println("++++++aaaa", err)
		BadRequest(gn, err)
		return
	}

	if len(request.Email) < 7 || !strings.Contains(request.Email, "@gmail.com") {
		BadRequest(gn, fmt.Errorf("email is not valid"))
		return
	}
	if len(request.Password) < 7 {
		BadRequest(gn, fmt.Errorf("password is not valid"))
		return
	}
	if len(request.Fullname) < 7 {
		BadRequest(gn, fmt.Errorf("fullname is not valid"))
		return
	}
	fmt.Println("++++++bbb", &request)
	_, err := h.AuthService.Register(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	Created(gn)
}

// LoginUser handles user login.
// @Summary Login User
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Create body genproto.LoginRequest true "Login"
// @Success 200 {object} string "Login Successful"
// @Failure 401 {string} string "Bad Request"
// @Failure 403 {string} string "Internal Server Error"
// @Router /api/user/login [post]
func (h *Handler) LoginUser(c *gin.Context) {
	var request pb.LoginRequest

	log.Printf("Received login request: %v", c.Request.Body)

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		BadRequest(c, err)
		return
	}

	if len(request.Email) < 7 || !strings.Contains(request.Email, "@") {
		BadRequest(c, fmt.Errorf("email is not valid"))
		return
	}
	if len(request.Password) < 7 {
		BadRequest(c, fmt.Errorf("password is not valid"))
		return
	}

	response, err := h.AuthService.Login(c, &request)
	if err != nil {
		log.Printf("Error in AuthService.Login: %v", err)
		InternalServerError(c, err)
		return
	}

	log.Printf("Login successful: %v", response)
	c.JSON(http.StatusOK, response)
}

// ResetPassword handles resetting a user's password.
// @Summary Reset Password
// @Description Reset a user's password
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.ResetPasswordRequest true "Reset Password"
// @Success 200 {string} string "Password Reset"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/reset-password [post]
func (h *Handler) ResetPassword(gn *gin.Context) {
	request := pb.ResetPasswordRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	response, err := h.AuthService.ResetPassword(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// RefreshToken handles refreshing a user's token.
// @Summary Refresh Token
// @Description Refresh a user's token
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} string "Token Refreshed"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/refresh-token [post]
func (h *Handler) RefreshToken(gn *gin.Context) {
	request := pb.RefreshTokenRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	response, err := h.AuthService.RefreshToken(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// Logout handles user logout.
// @Summary Logout
// @Description Logout a user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.LogoutRequest true "Logout"
// @Success 200 {object} string "User Logged Out"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/logout [post]
func (h *Handler) Logout(gn *gin.Context) {
	request := pb.LogoutRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	response, err := h.AuthService.Logout(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// GetProfile handles retrieving a user's profile.
// @Summary Get Profile
// @Description Get a user's profile
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} string "Profile Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/profile/{user_id} [get]
func (h *Handler) GetProfile(gn *gin.Context) {
	userID := gn.Param("user_id")
	if !IsValidUUID(userID) {
		BadRequest(gn, fmt.Errorf("user_id is not valid"))
		return
	}

	request := pb.GetProfileRequest{Userid: userID}
	response, err := h.AuthService.GetProfile(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// EditProfile handles editing a user's profile.
// @Summary Edit Profile
// @Description Edit a user's profile
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param Create body genproto.EditProfileRequest true "Edit Profile"
// @Success 200 {object} string "Profile Edited"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/profile/{user_id} [put]
func (h *Handler) EditProfile(gn *gin.Context) {
	userID := gn.Param("user_id")
	if !IsValidUUID(userID) {
		BadRequest(gn, fmt.Errorf("user_id is not valid"))
		return
	}

	request := pb.EditProfileRequest{Userid: userID}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	response, err := h.AuthService.EditProfile(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// ListUsers handles listing users.
// @Summary List Users
// @Description List users
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} string "Users Listed"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/users [get]
func (h *Handler) ListUsers(gn *gin.Context) {
	page, err := IsValidLimit(gn.Query("page"))
	if err != nil {
		BadRequest(gn, err)
		return
	}

	limit, err := IsValidLimit(gn.Query("limit"))
	if err != nil {
		BadRequest(gn, err)
		return
	}

	request := pb.ListUsersRequest{Page: int32(page), Limit: int32(limit)}
	response, err := h.AuthService.ListUsers(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// DeleteUser handles deleting a user.
// @Summary Delete User
// @Description Delete a user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Success 200 {string} string "User Deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/{user_id} [delete]
func (h *Handler) DeleteUser(gn *gin.Context) {
	userID := gn.Param("user_id")
	if !IsValidUUID(userID) {
		BadRequest(gn, fmt.Errorf("user_id is not valid"))
		return
	}

	request := pb.DeleteUserRequest{Userid: userID}
	response, err := h.AuthService.DeleteUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// GetEcoPoints handles retrieving eco points for a user.
// @Summary Get Eco Points
// @Description Get eco points for a user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} string "Eco Points Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/{user_id}/eco-points [get]
func (h *Handler) GetEcoPoints(gn *gin.Context) {
	userID := gn.Param("user_id")
	if userID == "" {
		BadRequest(gn, fmt.Errorf("user_id is required"))
		return
	}

	request := pb.GetEcoPointsRequest{Userid: userID}
	response, err := h.AuthService.GetEcoPoints(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// AddEcoPoints handles adding eco points for a user.
// @Summary Add Eco Points
// @Description Add eco points for a user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body genproto.AddEcoPointsRequest true "Add Eco Points"
// @Success 200 {object} string "Eco Points Added"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/eco-points [post]
func (h *Handler) AddEcoPoint(gn *gin.Context) {
	request := pb.AddEcoPointsRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	response, err := h.AuthService.AddEcoPoint(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// GetEcoPointsHistory handles retrieving eco points history for a user.
// @Summary Get Eco Points History
// @Description Get eco points history for a user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} string "Eco Points History Retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/{user_id}/eco-points/history [get]
func (h *Handler) GetEcoPointsHistory(gn *gin.Context) {
	userID := gn.Param("user_id")
	if userID == "" {
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

	request := pb.GetEcoPointsHistoryRequest{
		Userid: userID,
		Page:   int32(page),
		Limit:  int32(limit),
	}

	response, err := h.AuthService.GetEcoPointsHistory(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}
