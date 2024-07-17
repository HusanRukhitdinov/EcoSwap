package postgres

import (
	"database/sql"
	"fmt"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"log/slog"
	"time"
)

type SwapsRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewSwapsRepository(db *sql.DB) *SwapsRepository {
	return &SwapsRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *SwapsRepository) CreateChangeSwaps(request *pb.CreateChangeSwapRequest) (*pb.CreateChangeSwapResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO swaps (offered_item_id, requested_item_id, message, created_at) VALUES ($1, $2, $3, $4)", request.OfferedItemId, request.RequestedItemId, request.Message, time.Now())
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message create Swaps error -> %v", err))
		return nil, err
	}
	return &pb.CreateChangeSwapResponse{}, nil
}
func (repo *SwapsRepository) UpdateAcceptSwap(request *pb.UpdateAcceptSwapRequest) (*pb.UpdateAcceptSwapResponse, error) {
	_, err := repo.Db.Exec("UPDATE swaps SET status = $1, updated_at = $2 WHERE id = $3", "accepted", time.Now(), request.SwapId)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message update accept swap error -> %v", err))
		return nil, err
	}

	var response pb.UpdateAcceptSwapResponse
	err = repo.Db.QueryRow("SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, updated_at FROM swaps WHERE id = $1", request.SwapId).Scan(
		&response.Id, &response.OfferedItemId, &response.RequestedItemId, &response.RequesterId, &response.OwnerId, &response.Status, &response.UpdatedAt)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get updated swap error -> %v", err))
		return nil, err
	}

	return &response, nil
}
func (repo *SwapsRepository) UpdateRejectSwap(request *pb.UpdateRejactSwapRequest) (*pb.UpdateRejactSwapResponse, error) {
	_, err := repo.Db.Exec("UPDATE swaps SET status = $1, reason = $2, updated_at = $3 WHERE id = $4", "rejected", request.Reason, time.Now(), request.SwapId)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message update reject swap error -> %v", err))
		return nil, err
	}

	var response pb.UpdateRejactSwapResponse
	err = repo.Db.QueryRow("SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, reason, updated_at FROM swaps WHERE id = $1", request.SwapId).Scan(
		&response.Id, &response.OfferedItemId, &response.RequestedItemId, &response.RequesterId, &response.OwnerId, &response.Status, &response.Reason, &response.UpdatedAt)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get updated swap error -> %v", err))
		return nil, err
	}

	return &response, nil
}
func (repo *SwapsRepository) GetChangeSwap(request *pb.GetChangeSwapRequest) (*pb.GetChangeSwapResponse, error) {
	offset := (request.Page - 1) * request.Limit
	rows, err := repo.Db.Query("SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, created_at FROM swaps WHERE status = $1 LIMIT $2 OFFSET $3", request.Status, request.Limit, offset)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get change swap error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var swaps []*pb.Swaps
	for rows.Next() {
		var swap pb.Swaps
		err := rows.Scan(&swap.Id, &swap.OfferedItemId, &swap.RequestedItemId, &swap.RequesterId, &swap.OwnerId, &swap.Status, &swap.CreatedAt)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message scan swap error -> %v", err))
			return nil, err
		}
		swaps = append(swaps, &swap)
	}

	if err = rows.Err(); err != nil {
		repo.lg.Error(fmt.Sprintf("message rows error -> %v", err))
		return nil, err
	}

	var total int32
	err = repo.Db.QueryRow("SELECT COUNT(*) FROM swaps WHERE status = $1", request.Status).Scan(&total)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message count swaps error -> %v", err))
		return nil, err
	}

	response := &pb.GetChangeSwapResponse{
		Swaps: swaps,
		Total: total,
		Page:  request.Page,
		Limit: request.Limit,
	}

	return response, nil
}
