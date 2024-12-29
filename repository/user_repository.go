package repository

import (
	"github.com/paundraP/be-mcs/user-service/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByName(name string) bool {
	var user models.User
	if err := r.db.Where("name = ?", name).Find(&user).Error; err != nil {
		return false
	}
	if user.ID != "" {
		return true
	}
	return false
}

func (r *UserRepo) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}
