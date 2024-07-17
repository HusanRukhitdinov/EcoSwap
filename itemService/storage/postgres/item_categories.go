package postgres

import (
	"database/sql"
	"fmt"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
)

type ItemCategoryRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewItemCategoryRepository(db *sql.DB) *ItemCategoryRepository {
	return &ItemCategoryRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *ItemCategoryRepository) CreateItemCategory(request *pb.CreateItemCategoryManagRequest) (*pb.CreateItemCategoryManagResponse, error) {
	query := `INSERT INTO item_categories (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at`
	row := repo.Db.QueryRow(query, request.Name, request.Description)

	var response pb.CreateItemCategoryManagResponse
	err := row.Scan(&response.Id, &response.Name, &response.Description, &response.CreatedAt)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message add item category error -> %v", err))
		return nil, err
	}

	return &response, nil
}

func (repo *ItemCategoryRepository) GetStatistics(request *pb.GetStatisticsRequest) (*pb.GetStatisticsResponse, error) {
	startDate := request.StartDate
	endDate := request.EndDate

	var totalSwaps, totalRecycledItems, totalEcoPointsEarned int32

	err := repo.Db.QueryRow(`SELECT COUNT(*) FROM swaps WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalSwaps)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get total swaps error -> %v", err))
		return nil, err
	}

	err = repo.Db.QueryRow(`SELECT COUNT(*) FROM recycled_items WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalRecycledItems)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get total recycled items error -> %v", err))
		return nil, err
	}

	err = repo.Db.QueryRow(`SELECT SUM(eco_points) FROM eco_points WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalEcoPointsEarned)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get total eco points earned error -> %v", err))
		return nil, err
	}

	rows, err := repo.Db.Query(`SELECT id, name, COUNT(*) as swap_count FROM item_categories
	JOIN swaps ON item_categories.id = swaps.category_id WHERE swaps.created_at BETWEEN $1 AND $2
	GROUP BY id, name ORDER BY swap_count DESC LIMIT 10`, startDate, endDate)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get top categories error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var topCategories []*pb.TopCategories
	for rows.Next() {
		var category pb.TopCategories
		err := rows.Scan(&category.Id, &category.Name, &category.SwapCount)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message scan category error -> %v", err))
			return nil, err
		}
		topCategories = append(topCategories, &category)
	}

	rows, err = repo.Db.Query(`SELECT id, name, COUNT(*) as submissions_count FROM recycling_centers
	JOIN recycled_items ON recycling_centers.id = recycled_items.center_id WHERE recycled_items.created_at BETWEEN $1 AND $2
	GROUP BY id, name ORDER BY submissions_count DESC LIMIT 10`, startDate, endDate)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get top recycling centers error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var topRecyclingCenters []*pb.TopRecyclingCenters
	for rows.Next() {
		var center pb.TopRecyclingCenters
		err := rows.Scan(&center.Id, &center.Name, &center.SubmissionCount)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message scan center error -> %v", err))
			return nil, err
		}
		topRecyclingCenters = append(topRecyclingCenters, &center)
	}

	response := &pb.GetStatisticsResponse{
		TotalSwaps:           totalSwaps,
		TotalRecycledItems:   totalRecycledItems,
		TotalEcoPointsEarned: totalEcoPointsEarned,
		TopCategories:        topCategories,
		TopRecyclingCenters:  topRecyclingCenters,
	}

	return response, nil
}

func (repo *ItemCategoryRepository) GetMonitoringUserActivity(request *pb.GetMonitoringUserActivityRequest) (*pb.GetMonitoringUserActivityResponse, error) {

	var swapInitiated, swapCompleted, itemListed, recyclingSubmissions, ecoPointsEarned int32
	var userId string

	err := repo.Db.QueryRow(`SELECT COUNT(*) FROM swaps WHERE user_id = $1 AND created_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&swapInitiated)
	if err != nil {
		return nil, err
	}

	err = repo.Db.QueryRow(`SELECT COUNT(*) FROM swaps WHERE user_id = $1 AND completed_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&swapCompleted)
	if err != nil {
		return nil, err
	}

	err = repo.Db.QueryRow(`SELECT COUNT(*) FROM items WHERE user_id = $1 AND listed_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&itemListed)
	if err != nil {
		return nil, err
	}

	err = repo.Db.QueryRow(`SELECT COUNT(*) FROM recycled_items WHERE user_id = $1 AND submitted_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&recyclingSubmissions)
	if err != nil {
		return nil, err
	}

	err = repo.Db.QueryRow(`SELECT SUM(eco_points) FROM eco_points WHERE user_id = $1 AND earned_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&ecoPointsEarned)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMonitoringUserActivityResponse{
		UserId:               userId,
		SwapInitiated:        swapInitiated,
		SwapComplated:        swapCompleted,
		ItemListed:           itemListed,
		RecyclingSubmissions: recyclingSubmissions,
		EcoPointsEarned:      ecoPointsEarned,
	}

	return response, nil
}
