package controllers

import (
	"github.com/ahsansaif47/home-kitchens/auth/models"
	"github.com/ahsansaif47/home-kitchens/auth/repository/postgres"
)

type IUserService interface {
	// Define service methods here
}

type UserService struct {
	repo postgres.IUserRepository
}

func NewUserService(repo postgres.IUserRepository) IUserService {
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
