package postgres

import (
	"database/sql"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
)

type EcoTipsRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewEcoTipsRepository(db *sql.DB) *EcoTipsRepository {
	return &EcoTipsRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *EcoTipsRepository) CreateAddEcoTips(request *pb.CreateAddEcoTipsRequest) (*pb.CreateAddEcoTipsResponse, error) {
	var response pb.CreateAddEcoTipsResponse

	query := `INSERT INTO eco_tips (title, content, created_at) VALUES ($1, $2, NOW())
              RETURNING id, title, content, created_at;`

	err := repo.Db.QueryRow(query, request.Title, request.Content).
		Scan(&response.Id, &response.Title, &response.Content, &response.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (repo *EcoTipsRepository) GetAddEcoTips(request *pb.GetAddEcoTipsRequest) (*pb.GetAddEcoTipsResponse, error) {
	var response pb.GetAddEcoTipsResponse
	var tips []*pb.CreateAddEcoTipsResponse

	query := `SELECT id, title, content, created_at FROM eco_tips ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	rows, err := repo.Db.Query(query, request.Limit, (request.Page-1)*request.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tip pb.CreateAddEcoTipsResponse
		if err := rows.Scan(&tip.Id, &tip.Title, &tip.Content, &tip.CreatedAt); err != nil {
			return nil, err
		}
		tips = append(tips, &tip)
	}

	countQuery := `SELECT COUNT(*) FROM eco_tips;`
	var total int32
	err = repo.Db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, err
	}

	response.Tips = tips
	response.Total = total
	response.Page = request.Page
	response.Limit = request.Limit

	return &response, nil
}
