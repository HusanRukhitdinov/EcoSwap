package postgres

import (
	pb "eco_system/genproto"
	"fmt"
	"time"

	_ "github.com/form3tech-oss/jwt-go"
	"github.com/golang-jwt/jwt"
)

var secret_key = "salom"

type Token struct {
	AccessToken  string
	RefreshToken string
}

func (repo *UserRepository) GENERATEJWTToken(user *pb.LoginRequest) (*pb.LoginResponse, error) {
	AccessToken := jwt.New(jwt.SigningMethodHS256)
	claims := AccessToken.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["password"] = user.Password
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(3 * time.Hour).Unix()

	access, err := AccessToken.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println("Error signing access token")
		return nil, err
	}
	
	RefreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaim := RefreshToken.Claims.(jwt.MapClaims)
	refreshClaim["email"] = user.Email
	refreshClaim["password"] = user.Password
	refreshClaim["iat"] = time.Now().Unix()
	refreshClaim["exp"] = time.Now().Add(24 * time.Hour).Unix()
	
	refresh, err := RefreshToken.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println("Error signing refresh token")
		return nil, err
	}
	fmt.Println("---",&access,&refresh)
	return &pb.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		}, nil
	}
