package service

import (
	"context"
	pb "item_ser/genproto"
	"item_ser/storage/postgres"
)

type ItemService struct {
	ItemRepo                    *postgres.ItemRepository
	SwapRepo                    *postgres.SwapsRepository
	RecyclingCenterRepo         *postgres.RecyclingCentersRepository
	RecyclingSubmissionsRepo    *postgres.RecyclingSubmissionsRepository
	RatingsRepo                 *postgres.UserRatingRepository
	ItemCategoryRepo            *postgres.ItemCategoryRepository
	EcoChallengesRepo           *postgres.EcoChallengesRepository
	ChallengePartisipationsRepo *postgres.ChallengePartisipationsRepository
	EcoTipsRepo                 *postgres.EcoTipsRepository
	pb.UnimplementedEcoServiceServer
}

func NewItemService(
	item *postgres.ItemRepository,
	swap *postgres.SwapsRepository,
	recyclingCenter *postgres.RecyclingCentersRepository,
	recyclingSubmission *postgres.RecyclingSubmissionsRepository,
	ratings *postgres.UserRatingRepository,
	itemCategory *postgres.ItemCategoryRepository,
	ecoChallenges *postgres.EcoChallengesRepository,
	challengesPartisipatios *postgres.ChallengePartisipationsRepository,
	ecoTips *postgres.EcoTipsRepository,
) *ItemService {
	return &ItemService{
		ItemRepo:                    item,
		SwapRepo:                    swap,
		RecyclingCenterRepo:         recyclingCenter,
		RecyclingSubmissionsRepo:    recyclingSubmission,
		RatingsRepo:                 ratings,
		ItemCategoryRepo:            itemCategory,
		EcoChallengesRepo:           ecoChallenges,
		ChallengePartisipationsRepo: challengesPartisipatios,
		EcoTipsRepo:                 ecoTips,
	}
}

func (service *ItemService) CreateItem(ctx context.Context, in *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	return service.ItemRepo.CreateItem(in)
}
func (service *ItemService) UpdateItem(ctx context.Context,in *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	return service.ItemRepo.UpdateItem(in)
}
func (service *ItemService) DeleteItem(ctx context.Context, in *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	return service.ItemRepo.DeleteItem(in)
}
func (service *ItemService) GetAllItems(ctx context.Context, in *pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error) {
	return service.ItemRepo.GetAllItems(in)
}
func (service *ItemService) GetByIdItem(ctx context.Context, in *pb.GetByIdItemRequest) (*pb.GetByIdItemResponse, error) {
	return service.ItemRepo.GetByIdItem(in)
}
func (service *ItemService) SearchItemsAndFilter(ctx context.Context, in *pb.SearchItemsAndFiltRequest) (*pb.SearchItemsAndFiltResponse, error) {
	return service.ItemRepo.SearchItemsAndFilter(in)
}


func (service *ItemService) CreateChangeSwaps(ctx context.Context, in *pb.CreateChangeSwapRequest) (*pb.CreateChangeSwapResponse, error) {
	return service.SwapRepo.CreateChangeSwaps(in)
}
func (service *ItemService) UpdateAcceptSwap(ctx context.Context, in *pb.UpdateAcceptSwapRequest) (*pb.UpdateAcceptSwapResponse, error) {
	return service.SwapRepo.UpdateAcceptSwap(in)
}
func (service *ItemService) UpdateRejectSwap(ctx context.Context, in *pb.UpdateRejactSwapRequest) (*pb.UpdateRejactSwapResponse, error) {
	return service.SwapRepo.UpdateRejectSwap(in)
}
func (service *ItemService) GetChangeSwap(ctx context.Context, in *pb.GetChangeSwapRequest) (*pb.GetChangeSwapResponse, error) {
	return service.SwapRepo.GetChangeSwap(in)
}


func (service *ItemService) CreateAddRecyclingCenter(ctx context.Context, in *pb.CreateAddRecyclingCenterRequest) (*pb.CreateAddRecyclingCenterResponse, error) {
	return service.RecyclingCenterRepo.CreateAddRecyclingCenter(in)
}
func (service *ItemService) SearchRecyclingCenter(ctx context.Context, in *pb.SearchRecyclingCenterRequest) (*pb.SearchRecyclingCenterResponse, error) {
	return service.RecyclingCenterRepo.SearchRecyclingCenter(in)
}


func (service *ItemService) CreateRecyclingSubmission(ctx context.Context, in *pb.CreteRecyclingSubmissionsRequest) (*pb.CreteRecyclingSubmissionsResponse, error) {
	return service.RecyclingSubmissionsRepo.CreateRecyclingSubmission(in)
}


func (service *ItemService) CreateAddUserRating(ctx context.Context, in *pb.CreateAddUserRatingRequest) (*pb.CreateAddUserRatingResponse, error) {
	return service.RatingsRepo.CreateAddUserRating(in)
}
func (service *ItemService) GetUserRatings(ctx context.Context, in *pb.GetUserRatingRequest) (*pb.GetUserRatingResponse, error) {
	return service.RatingsRepo.GetUserRatings(in)
}


func (service *ItemService) CreateItemCategory(ctx context.Context, in *pb.CreateItemCategoryManagRequest) (*pb.CreateItemCategoryManagResponse, error) {
	return service.ItemCategoryRepo.CreateItemCategory(in)
}
func (service *ItemService) GetStatistics(ctx context.Context, in *pb.GetStatisticsRequest) (*pb.GetStatisticsResponse, error) {
	return service.ItemCategoryRepo.GetStatistics(in)
}
func (service *ItemService) GetMonitoringUserActivity(ctx context.Context, in *pb.GetMonitoringUserActivityRequest) (*pb.GetMonitoringUserActivityResponse, error) {
	return service.ItemCategoryRepo.GetMonitoringUserActivity(in)
}


func (service *ItemService) CreateEcoChallenge(ctx context.Context, in *pb.CreateEcoChallengeRequest) (*pb.CreateEcoChallengeResponse, error) {
	return service.EcoChallengesRepo.CreateEcoChallenge(in)
}


func (service *ItemService) ParticipateChallenge(ctx context.Context, in *pb.CreateParticipateChallengeRequest) (*pb.CreateParticipateChallengeResponse, error) {
	return service.ChallengePartisipationsRepo.ParticipateChallenge(in)
}
func (service *ItemService) UpdateEcoChallengeResult(ctx context.Context, in *pb.UpdateEcoChallengeRresultRequest) (*pb.UpdateEcoChallengeRresultResponse, error) {
	return service.ChallengePartisipationsRepo.UpdateEcoChallengeResult(in)
}


func (service *ItemService) CreateAddEcoTips(ctx context.Context, in *pb.CreateAddEcoTipsRequest) (*pb.CreateAddEcoTipsResponse, error) {
	return service.EcoTipsRepo.CreateAddEcoTips(in)
}
func (service *ItemService) GetAddEcoTips(ctx context.Context, in *pb.GetAddEcoTipsRequest) (*pb.GetAddEcoTipsResponse, error) {
	return service.EcoTipsRepo.GetAddEcoTips(in)
}