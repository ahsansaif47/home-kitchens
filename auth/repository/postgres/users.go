package postgres

import (
	"errors"

	"github.com/ahsansaif47/home-kitchens/auth/constants"
	"github.com/ahsansaif47/home-kitchens/auth/models"
	"github.com/ahsansaif47/home-kitchens/auth/utils"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CheckExistingEmail(email string) (bool, error)
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id string) (*models.User, error)
	GetAllVendors() ([]models.User, error)
	GetAllUsers() ([]models.User, error)
	SetNewPassword(email, newPassword string) (bool, error)
	ValidateUserCredentials(email, password string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CheckExistingEmail(email string) (bool, error) {
	err := r.db.Where("email = ?", email).First(&models.User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	status, err := r.CheckExistingEmail(user.Email)
	if err != nil {
		return err
	}
	if status {
		return constants.ErrUserAlreadyExists
	}

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

func (r *UserRepository) SetNewPassword(email, newPassword string) (bool, error) {
	passwordHash, err := utils.GeneratePasswordHash(newPassword)
	if err != nil {
		return false, err
	}
	if err := r.db.Model(&models.User{}).Where("email = ?", email).Update("password_hash", passwordHash).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) ValidateUserCredentials(email, password string) (*models.User, error) {
	var user models.User

	err := r.db.Where("LOWER(email) = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(password, *user.PasswordHash) {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
