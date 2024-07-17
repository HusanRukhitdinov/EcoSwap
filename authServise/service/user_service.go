package service

import (
	"context"
	pb "eco_system/genproto"
	"eco_system/storage/postgres"
)

type UserService struct {
	UserRepo *postgres.UserRepository
	pb.UnimplementedAuthServiceServer
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (service *UserService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return service.UserRepo.Register(in)
}
func (service *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return service.UserRepo.Login(in)
}
func (service *UserService) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	return service.UserRepo.GetProfile(in)
}
func (service *UserService) EditProfile(ctx context.Context, in *pb.EditProfileRequest) (*pb.EditProfileResponse, error) {
	return service.UserRepo.EditProfile(in)
}
func (service *UserService) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return service.UserRepo.ListUsers(in)
}
func (service *UserService) DeleteUser(ctx context.Context,in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return service.UserRepo.DeleteUser(in)
}
func (service *UserService) ResetPassword(ctx context.Context, in *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	return service.UserRepo.ResetPassword(in)
}
func (service *UserService) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return service.UserRepo.RefreshToken(in)
}
func (service *UserService) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return service.UserRepo.Logout(in)
}
func (service *UserService) GetEcoPoint(ctx context.Context, in *pb.GetEcoPointsRequest) (*pb.GetEcoPointsResponse, error) {
	return service.UserRepo.GetEcoPoint(in)
}
func (service *UserService) AddEcoPoint(ctx context.Context, in *pb.AddEcoPointsRequest) (*pb.AddEcoPointsResponse, error) {
	return service.UserRepo.AddEcoPoint(in)
}
func (service *UserService) GetEcoPointsHistory(ctx context.Context, in *pb.GetEcoPointsHistoryRequest) (*pb.GetEcoPointsHistoryResponse, error) {
	return service.UserRepo.GetEcoPointsHistory(in)
}
func (service *UserService) GENERATEJWTToken(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	jwtToken, err := service.UserRepo.GENERATEJWTToken(in)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  jwtToken.AccessToken,
		RefreshToken: jwtToken.RefreshToken,
		ExpiresIn:    jwtToken.ExpiresIn,
	}, nil
}
