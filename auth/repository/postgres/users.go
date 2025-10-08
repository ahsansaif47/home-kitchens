package postgres

import (
	"github.com/ahsansaif47/home-kitchens/auth/constants"
	"github.com/ahsansaif47/home-kitchens/auth/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id string) (*models.User, error)
	GetAllVendors() ([]models.User, error)
	GetAllUsers() ([]models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// Methods for IUserRepository here
func (r *UserRepository) CreateUser(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", id)
	return &user, result.Error
}

func (r *UserRepository) GetAllVendors() ([]models.User, error) {
	var users []models.User
	result := r.db.Where("role_id = ?", constants.RoleVendor).Find(&users)
	return users, result.Error
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.db.Where("role_id = ?", constants.RoleUser).Find(&users)
	return users, result.Error
}
