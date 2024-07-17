package postgres

import (
	"database/sql"
	"fmt"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
	"time"
)

type RecyclingSubmissionsRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewRecyclingSubmissionsRepository(db *sql.DB) *RecyclingSubmissionsRepository {
	return &RecyclingSubmissionsRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *RecyclingSubmissionsRepository) CreateRecyclingSubmission(request *pb.CreteRecyclingSubmissionsRequest) (*pb.CreteRecyclingSubmissionsResponse, error) {
	tx, err := repo.Db.Begin()
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message begin transaction error -> %v", err))
		return nil, err
	}

	query := `INSERT INTO recycling_submissions (center_id, user_id, eco_point_earned, created_at) 
	          VALUES ($1, $2, $3, $4) RETURNING id, center_id, user_id, eco_point_earned, created_at`
	row := tx.QueryRow(query, request.CenterId, request.UserId, calculateEcoPoints(request.Items),time.Now())

	var response pb.CreteRecyclingSubmissionsResponse
	err = row.Scan(&response.Id, &response.CenterId, &response.UserId, &response.EcoPointEarned, &response.CreatedAt)
	if err != nil {
		tx.Rollback()
		repo.lg.Error(fmt.Sprintf("message create recycling submission error -> %v", err))
		return nil, err
	}

	for _, item := range request.Items {
		itemQuery := `INSERT INTO recycling_items (submission_id, item_id, weight, material) VALUES ($1, $2, $3, $4)`
		_, err := tx.Exec(itemQuery, response.Id, item.ItemId, item.Weight, item.Material)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message insert item error -> %v", err))
			return nil, err
		}
	}
	response.Items = request.Items
	return &response, nil
}

func calculateEcoPoints(items []*pb.Itemes) int32 {
	var totalPoints int32
	for _, item := range items {
		totalPoints += int32(item.Weight * 10) 
	}
	return totalPoints
}
