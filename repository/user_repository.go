package repository

import (
	"github.com/paundraP/be-mcs/user-service/graphql/generated"
	gn "github.com/paundraP/be-mcs/user-service/graphql/generated"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserByID(id string) (*gn.User, error) {
	var user gn.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return &gn.User{}, err
	}

	return &user, nil
}

func (r *UserRepo) CheckEmail(email string) bool {
	if err := r.db.Where("email = ?", email).Error; err != nil {
		return false
	}
	return true
}

func (r *UserRepo) GetUserByEmail(email string) (*generated.User, error) {
	var user gn.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetAllUsers() ([]*gn.User, error) {
	var users []*gn.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) CreateUser(user *gn.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&gn.User{}).Error
}
