package postgres

import (
	"database/sql"
	"fmt"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
	"time"
)

type RecyclingCentersRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewRecyclingCentersRepository(db *sql.DB) *RecyclingCentersRepository {
	return &RecyclingCentersRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *RecyclingCentersRepository) CreateAddRecyclingCenter(request *pb.CreateAddRecyclingCenterRequest) (*pb.CreateAddRecyclingCenterResponse, error) {
	query := `INSERT INTO recycling_centers (name, addres, accepted_materials, working_hours, contact_number, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, addres, accepted_materials, working_hours, contact_number, created_at`
	row := repo.Db.QueryRow(query, request.Name, request.Addres, request.AcceptedMaterials, request.WorkingHours, request.ContactNumber,time.Now())

	var response pb.CreateAddRecyclingCenterResponse
	err := row.Scan(&response.Id, &response.Name, &response.Addres, &response.AcceptedMaterials, &response.WorkingHours, &response.ContactNumber, &response.CreatedAt)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message create add recycling center error -> %v", err))
		return nil, err
	}

	return &response, nil
}

func (repo *RecyclingCentersRepository) SearchRecyclingCenter(request *pb.SearchRecyclingCenterRequest) (*pb.SearchRecyclingCenterResponse, error) {
	offset := (request.Page - 1) * request.Limit
	rows, err := repo.Db.Query("SELECT id, name, addres, accepted_materials, working_hours, contact_number FROM recycling_centers WHERE accepted_materials ILIKE $1 LIMIT $2 OFFSET $3", "%"+request.Material+"%", request.Limit, offset)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message search recycling center error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var centers []*pb.Centers
	for rows.Next() {
		var center pb.Centers
		err := rows.Scan(&center.Id, &center.Name, &center.Addres, &center.AcceptedMaterials, &center.WorkingHours, &center.ContactNumber)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message scan center error -> %v", err))
			return nil, err
		}
		centers = append(centers, &center)
	}

	if err = rows.Err(); err != nil {
		repo.lg.Error(fmt.Sprintf("message rows error -> %v", err))
		return nil, err
	}

	var total int32
	err = repo.Db.QueryRow("SELECT COUNT(*) FROM recycling_centers WHERE accepted_materials ILIKE $1", "%"+request.Material+"%").Scan(&total)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message count centers error -> %v", err))
		return nil, err
	}

	response := &pb.SearchRecyclingCenterResponse{
		Centers: centers,
		Total:   total,
		Page:    request.Page,
		Limit:   request.Limit,
	}

	return response, nil
}
