package postgres

import (
	"database/sql"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
	"time"
)

type ChallengePartisipationsRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewChallengePrtisipationsRepository(db *sql.DB) *ChallengePartisipationsRepository {
	return &ChallengePartisipationsRepository{Db: db, lg: logger.NewLogger()}
}
func (repo *ChallengePartisipationsRepository) ParticipateChallenge(request *pb.CreateParticipateChallengeRequest) (*pb.CreateParticipateChallengeResponse, error) {
    var response pb.CreateParticipateChallengeResponse

    query := `INSERT INTO eco_challenge_participants (challenge_id,status, joined_at)
              VALUES ($1, $2, $3)
              RETURNING challenge_id, user_id, status, joined_at;`

    err := repo.Db.QueryRow(query, request.ChallengeId, time.Now().UTC()).
	Scan(&response.ChallengeId, &response.UserId, &response.Status, &response.JoinedAt)
    
    if err != nil {
        return nil, err
    }

    return &response, nil
}

func (repo *ChallengePartisipationsRepository) UpdateEcoChallengeResult(request *pb.UpdateEcoChallengeRresultRequest) (*pb.UpdateEcoChallengeRresultResponse, error) {
    var response pb.UpdateEcoChallengeRresultResponse

    query := `UPDATE eco_challenge_participants SET recycled_items_count = $1,challenge_id = $2 updated_at = $3
              WHERE challenge_id = $3,user_id = $4, deleted_at is null
              RETURNING challenge_id, user_id, status, recycled_items_count, updated_at;`

    err := repo.Db.QueryRow(query, request.RecycledItemsCount,request.ChallengeId, time.Now().UTC(), request.ChallengeId, request.ChallengeId).
	Scan(&response.ChallengeId, &response.UserId, &response.Status, &response.RecycledItemsCount, &response.UpdatedAt)
    
    if err != nil {
        return nil, err
    }

    return &response, nil
}
