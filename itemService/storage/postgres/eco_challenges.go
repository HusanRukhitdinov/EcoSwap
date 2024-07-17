package postgres

import (
	"database/sql"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
	"time"
)

type EcoChallengesRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewEcoChallengesRepository(db *sql.DB) *EcoChallengesRepository {
	return &EcoChallengesRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *EcoChallengesRepository) CreateEcoChallenge(request *pb.CreateEcoChallengeRequest) (*pb.CreateEcoChallengeResponse, error) {
	var challenge pb.CreateEcoChallengeResponse

	query := `INSERT INTO eco_challenges (title, description, start_date, end_date, reward_points, created_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7)
              RETURNING id, title, description, start_date, end_date, reward_points, created_at;`

	err := repo.Db.QueryRow(query, request.Title, request.Description, request.StartDate, request.EndDate, request.RewardPoints, time.Now().UTC()).
		Scan(&challenge.Id, &challenge.Title, &challenge.Description, &challenge.StartDate, &challenge.EndDate, &challenge.RewardPoints, &challenge.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &challenge, nil
}
