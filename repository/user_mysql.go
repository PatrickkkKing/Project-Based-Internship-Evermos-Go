package repository

import (
	"evermos/domain"

	"gorm.io/gorm"
)

type userMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserRepository {
	return &userMySQL{db}
}

func (r *userMySQL) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userMySQL) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userMySQL) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userMySQL) Update(user *domain.User) error {
	return r.db.Save(user).Error
}
