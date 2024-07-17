package postgres

import (
	"database/sql"
	"eco_system/config/logger"
	pb "eco_system/genproto"
	storage "eco_system/help"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	Db    *sql.DB
	Loger *slog.Logger
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db, Loger: logger.NewLogger()}
}

func (repo *UserRepository) Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id := uuid.NewString()
	fmt.Println("sc++++")
	_, err := repo.Db.Exec(`
	INSERT INTO users(id, email, password, fullname, created_at)
	VALUES($1, $2, $3, $4, $5) `,
		id, request.Email, request.Password, request.Fullname, time.Now())
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in Register: %v", err)
		return nil, err
	}
	return &pb.RegisterResponse{
		Id:        id,
		Email:     request.Email,
		Fullname:  request.Fullname,
		CreatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (repo *UserRepository) Login(request *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokenResponse, err := repo.GENERATEJWTToken(request)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error generate JWT token: %v", err)
		return nil, err
	}

	_, err = repo.Db.Exec("UPDATE users SET refresh_token=$1 WHERE email=$2 AND password=$3", tokenResponse.RefreshToken, request.Email, request.Password)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error update refresh token: %v", err)
		return nil, err
	}

	return tokenResponse, nil
}

func (repo *UserRepository) GetProfile(request *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	var profResponse pb.GetProfileResponse

	query := "SELECT id, email, fullname, ecopoints, created_at, updated_at FROM users WHERE id = $1"
	err := repo.Db.QueryRow(query, request.Userid).Scan(&profResponse.Id, &profResponse.Email, &profResponse.Fullname, &profResponse.EcoPoints, &profResponse.CreatedAt, &profResponse.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with id: %v", request.Userid)
			return nil, fmt.Errorf("no user found with id: %v", request.Userid)
		}
		log.Printf("Error in GetProfile: %v", err)
		return nil, fmt.Errorf("error retrieving profile: %v", err)
	}

	return &profResponse, nil
}

func (repo *UserRepository) EditProfile(request *pb.EditProfileRequest) (*pb.EditProfileResponse, error) {
	_, err := repo.Db.Exec("UPDATE users SET fullname = $1, bio = $2, updated_at = $3 WHERE id = $4", request.Fullname, request.Bio, time.Now(), request.Userid)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in EditProfile: %v", err)
		return nil, err
	}
	return &pb.EditProfileResponse{
		Id:        request.Userid,
		Fullname:  request.Fullname,
		Bio:       request.Bio,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (repo *UserRepository) ListUsers(request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	filter := ""

	if request.Page > 0 {
		params["page"] = request.Page
		filter += " OFFSET :page"
	}
	if request.Limit > 0 {
		params["limit"] = request.Limit
		filter += " LIMIT :limit"
	}

	query := "SELECT id, fullname, ecopoints FROM users WHERE deleted_at IS NULL" + filter
	query, arr = storage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in ListUsers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var userResponse pb.User
		err := rows.Scan(&userResponse.Id, &userResponse.Fullname, &userResponse.EcoPoints)
		if err != nil {
			repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
			log.Printf("Error in ListUsers (row scan): %v", err)
			return nil, err
		}
		users = append(users, &userResponse)
	}
	return &pb.ListUsersResponse{
		Users: users,
		Total: int32(len(users)),
		Page:  request.Page,
		Limit: request.Limit,
	}, nil
}

func (repo *UserRepository) DeleteUser(request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := repo.Db.Exec("UPDATE users SET deleted_at = current_timestamp WHERE id = $1 AND deleted_at IS NULL", request.Userid)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in DeleteUser: %v", err)
		return nil, err
	}
	return &pb.DeleteUserResponse{Message: "User deleted successfully"}, nil
}

func (repo *UserRepository) ResetPassword(request *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	_, err := repo.Db.Exec("UPDATE users SET password= $1 WHERE email = $2 AND deleted_at IS NULL", "new_password", request.Email)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in ResetPassword: %v", err)
		return nil, err
	}
	return &pb.ResetPasswordResponse{Message: "Password reset successfully"}, nil
}

func (repo *UserRepository) RefreshToken(request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO users (refresh_token) VALUES ($1)", request.RefreshToken)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in RefreshToken: %v", err)
		return nil, err
	}
	return &pb.RefreshTokenResponse{
		AccessToken:  "new_access_token",
		RefreshToken: request.RefreshToken,
	}, nil
}

func (repo *UserRepository) Logout(request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return &pb.LogoutResponse{Message: "User logged out successfully"}, nil
}

func (repo *UserRepository) GetEcoPoints(request *pb.GetEcoPointsRequest) (*pb.GetEcoPointsResponse, error) {
	var ecoPointsResponse pb.GetEcoPointsResponse
	err := repo.Db.QueryRow("SELECT user_id, ecopoints FROM users WHERE user_id = $1", request.Userid).Scan(&ecoPointsResponse.Userid, &ecoPointsResponse.EcoPoints)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in GetEcoPoint: %v", err)
		return nil, err
	}
	return &ecoPointsResponse, nil
}

func (repo *UserRepository) AddEcoPoint(request *pb.AddEcoPointsRequest) (*pb.AddEcoPointsResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO users(user_id, ecopoints, reason, timestamp) VALUES ($1, $2, $3, $4)", request.Userid, request.Points, request.Reason, time.Now())
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in AddEcoPoint: %v", err)
		return nil, err
	}
	return &pb.AddEcoPointsResponse{
		Userid:      request.Userid,
		EcoPoints:   request.Points,
		AddedPoints: request.Points,
		Reason:      request.Reason,
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}

func (repo *UserRepository) GetEcoPointsHistory(request *pb.GetEcoPointsHistoryRequest) (*pb.GetEcoPointsHistoryResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	filter := ""

	if request.Page > 0 {
		params["page"] = request.Page
		filter += " OFFSET :page"
	}
	if request.Limit > 0 {
		params["limit"] = request.Limit
		filter += " LIMIT :limit"
	}

	query := "SELECT id, points, type, reason, timestamp FROM eco_points_history WHERE user_id = $1" + filter
	query, arr = storage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
		log.Printf("Error in GetEcoPointsHistory: %v", err)
		return nil, err
	}
	defer rows.Close()

	var history []*pb.EcoPointsHistory
	for rows.Next() {
		var ecoPointsHistory pb.EcoPointsHistory
		err := rows.Scan(&ecoPointsHistory.Id, &ecoPointsHistory.Points, &ecoPointsHistory.Type, &ecoPointsHistory.Reason, &ecoPointsHistory.Timestamp)
		if err != nil {
			repo.Loger.Error(fmt.Sprintf("error getting data: %v",err))
			log.Printf("Error in GetEcoPointsHistory (row scan): %v", err)
			return nil, err
		}
		history = append(history, &ecoPointsHistory)
	}
	return &pb.GetEcoPointsHistoryResponse{
		History: history,
		Total:   int32(len(history)),
		Page:    request.Page,
		Limit:   request.Limit,
	}, nil
}
