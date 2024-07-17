package postgres

import (
	"database/sql"
	"fmt"
	"item_ser/config/logger"
	pb "item_ser/genproto"
	"item_ser/help"
	"log/slog"
	"time"
)

type ItemRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *ItemRepository) CreateItem(request *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO items (name, description, category_id, condition, swap_preference, images, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
	 request.Name, request.Description, request.CategoryId, request.Condition, request.SwapPreference, request.Images, time.Now())
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message create item error -> %v", err))
		return nil, err
	}
	return &pb.CreateItemResponse{}, nil
}

func (repo *ItemRepository) UpdateItem(request *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	_, err := repo.Db.Exec("UPDATE items SET name = $1, condition = $2, updated_at = $3 WHERE id = $4 AND deleted_at IS NULL", request.Name, request.Condition, time.Now(),request.Id)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message update item error -> %v", err))
		return nil, err
	}
	return &pb.UpdateItemResponse{}, nil
}

func (repo *ItemRepository) DeleteItem(request *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	_, err := repo.Db.Exec("UPDATE items SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL", time.Now(), request.ItemId)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message delete item error -> %v", err))
		return nil, err
	}
	return &pb.DeleteItemResponse{}, nil
}

func (repo *ItemRepository) GetAllItems(request *pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	filter := " WHERE deleted_at IS NULL"
	if len(request.Id) > 0 {
		params["id"] = request.Id
		filter += " AND id = :id"
	}
	if len(request.Name) > 0 {
		params["name"] = request.Name
		filter += " AND name = :name"
	}
	if len(request.CategoryId) > 0 {
		params["category_id"] = request.CategoryId
		filter += " AND category_id = :category_id"
	}
	if len(request.Condition) > 0 {
		params["condition"] = request.Condition
		filter += " AND condition = :condition"
	}
	if len(request.OwnerId) > 0 {
		params["owner_id"] = request.OwnerId
		filter += " AND owner_id = :owner_id"
	}
	if len(request.Status) > 0 {
		params["status"] = request.Status
		filter += " AND status = :status"
	}
	if request.Limit > 0 {
		params["limit"] = request.Limit
		filter += " LIMIT :limit"
	}
	if request.Offset > 0 {
		params["offset"] = request.Offset
		filter += " OFFSET :offset"
	}
	query := "SELECT id, name, category_id, condition, owner_id, status FROM items" + filter
	query, arr = help.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get all items error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var items []*pb.Items
	for rows.Next() {
		var item pb.Items
		err := rows.Scan(&item.Id, &item.Name, &item.CategoryId, &item.Condition, &item.OwnerId, &item.Status)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message get all items scan error -> %v", err))
			return nil, err
		}
		items = append(items, &item)
	}
	return &pb.GetAllItemsResponse{Items: items}, nil
}

func (repo *ItemRepository) GetByIdItem(request *pb.GetByIdItemRequest) (*pb.GetByIdItemResponse, error) {
	var response pb.GetByIdItemResponse
	err := repo.Db.QueryRow("SELECT id, name, description, category_id, condition, swap_preference, owner_id, status, created_at, updated_at FROM items WHERE deleted_at IS NULL AND id = $1", request.ItemId).
	Scan(&response.Id, &response.Name, &response.Description, &response.CategoryId, &response.Condition, &response.SwapPreference, &response.OwnerId, &response.Status, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message get by id item error -> %v", err))
		return nil, err
	}
	return &response, nil
}

func (repo *ItemRepository) SearchItemsAndFilter(request *pb.SearchItemsAndFiltRequest) (*pb.SearchItemsAndFiltResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	filter := " WHERE deleted_at IS NULL"
	if len(request.Query) > 0 {
		params["query"] = "%" + request.Query + "%"
		filter += " AND (name ILIKE :query OR description ILIKE :query)"
	}
	if len(request.Category) > 0 {
		params["category"] = request.Category
		filter += " AND category_id = :category"
	}
	if len(request.Condition) > 0 {
		params["condition"] = request.Condition
		filter += " AND condition = :condition"
	}
	if request.Page > 0 {
		params["page"] = request.Page
		filter += " LIMIT :limit OFFSET :offset"
		params["limit"] = request.Limit
		params["offset"] = (request.Page - 1) * request.Limit
	}
	query := "SELECT id, name, category_id, condition, owner_id, status FROM items" + filter
	query, arr = help.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		repo.lg.Error(fmt.Sprintf("message search items and filter error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var items []*pb.Items
	for rows.Next() {
		var item pb.Items
		err := rows.Scan(&item.Id, &item.Name, &item.CategoryId, &item.Condition, &item.OwnerId, &item.Status)
		if err != nil {
			repo.lg.Error(fmt.Sprintf("message search items and filter scan error -> %v", err))
			return nil, err
		}
		items = append(items, &item)
	}
	return &pb.SearchItemsAndFiltResponse{Items: items}, nil
}
