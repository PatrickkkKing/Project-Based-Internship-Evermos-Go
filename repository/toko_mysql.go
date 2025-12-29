package repository

import (
	"evermos/domain"

	"gorm.io/gorm"
)

type tokoMySQL struct {
	db *gorm.DB
}

func NewTokoMySQL(db *gorm.DB) TokoRepository {
	return &tokoMySQL{db}
}

func (r *tokoMySQL) Create(toko *domain.Toko) error {
	return r.db.Create(toko).Error
}

func (r *tokoMySQL) FindByUserID(userID uint) (*domain.Toko, error) {
	var toko domain.Toko
	err := r.db.Where("user_id = ?", userID).First(&toko).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoMySQL) Update(toko *domain.Toko) error {
	return r.db.Save(toko).Error
}
