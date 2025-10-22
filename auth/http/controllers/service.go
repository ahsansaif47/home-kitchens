package controllers

import (
	"context"
	"time"

	"github.com/ahsansaif47/home-kitchens/auth/constants"
	"github.com/ahsansaif47/home-kitchens/auth/gRPC/services"
	ec "github.com/ahsansaif47/home-kitchens/auth/gRPC/services"
	"github.com/ahsansaif47/home-kitchens/auth/models"
	"github.com/ahsansaif47/home-kitchens/auth/repository/postgres"
	"github.com/ahsansaif47/home-kitchens/auth/repository/redis"
	"github.com/ahsansaif47/home-kitchens/auth/utils"
)

type IUserService interface {
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id string) (*models.User, error)
	GetAllVendors() ([]models.User, error)
	GetAllUsers() ([]models.User, error)
	SetNewPassword(email, newPassword string) (bool, error)
}

type UserService struct {
	repo        postgres.IUserRepository
	cacheRepo   redis.ICacheRepository
	emailClient ec.EmailClient
}

func NewUserService(repo postgres.IUserRepository, cache redis.ICacheRepository, ec services.EmailClient) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) FindAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) FindByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllVendors() ([]models.User, error) {
	return s.repo.GetAllVendors()
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) SetNewPassword(email, newPassword string) (bool, error) {
	return s.repo.SetNewPassword(email, newPassword)
}

func (s *UserService) ValidateUserCredentials(email, password string) (*models.User, error) {
	return s.repo.ValidateUserCredentials(email, password)
}

func (s *UserService) GenerateAndSendOTP(email string) error {
	otp := utils.GenerateOTP()
	// TODO:
	// 1. Store OTP in Redis with expiration
	// 2. Send OTP to user's email

	otpHash := utils.HashOTP(otp)
	err := s.cacheRepo.StoreOTP(email, otpHash, 1*time.Minute)
	if err != nil {
		return err
	}
	// Send the OTP to user's email
	// RPC endpoint to send the OTP to emailing service

	// TODO: Pick this email from notifications service later on
	err = s.emailClient.SendOTPEmail(context.Background(), constants.HomeKitchensEmail, email, "OTP Verification Email", "", otp)
	if err != nil {
		return err
	}

	return nil
}
