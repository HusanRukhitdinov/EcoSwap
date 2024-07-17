package postgres

import (
	"database/sql"
	"fmt"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
	"time"
)

type UserRatingRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewUserRatingRepository(db *sql.DB) *UserRatingRepository {
	return &UserRatingRepository{Db: db, lg: logger.NewLogger()}
}
func (repo *UserRatingRepository) CreateAddUserRating(request *pb.CreateAddUserRatingRequest) (*pb.CreateAddUserRatingResponse, error) {
	query := `INSERT INTO user_ratings (user_id, rating, comment, swap_id) 
	          VALUES ($1, $2, $3, $4) RETURNING id, user_id, rater_id, rating, comment, swap_id, created_at`
	row := repo.Db.QueryRow(query, request.UserId, request.Rating, request.Comment, request.SwapId,time.Now())

	var response pb.CreateAddUserRatingResponse
	err := row.Scan(&response.Id, &response.UserId, &response.RaterId, &response.Rating, &response.Comment, &response.SwapId, &response.CreatedAt)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message add user rating error -> %v", err))
		return nil, err
	}

	return &response, nil
}

func (repo *UserRatingRepository) GetUserRatings(request *pb.GetUserRatingRequest) (*pb.GetUserRatingResponse, error) {
	limit := request.Limit
	offset := (request.Page - 1) * limit

	query := `SELECT id, rater_id, rating, comment, swap_id, created_at FROM user_ratings WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := repo.Db.Query(query, request.UserId, limit, offset)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get user ratings error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var ratings []*pb.Rating
	for rows.Next() {
		var rating pb.Rating
		err := rows.Scan(&rating.Id, &rating.RaterId, &rating.Rating, &rating.Comment, &rating.SwapId, &rating.CreatedAt)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message scan rating error -> %v", err))
			return nil, err
		}
		ratings = append(ratings, &rating)
	}

	avgQuery := `SELECT AVG(rating) as average_rating, COUNT(*) as total_ratings FROM user_ratings WHERE user_id = $1`
	var averageRating float64
	var totalRatings int32
	err = repo.Db.QueryRow(avgQuery, request.UserId).Scan(&averageRating, &totalRatings)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get average rating error -> %v", err))
		return nil, err
	}

	response := &pb.GetUserRatingResponse{
		Ratings:       ratings,
		AverageRating: float32(averageRating),
		TotalRatings:  totalRatings,
		Page:          request.Page,
		Limit:         limit,
	}

	return response, nil
}
