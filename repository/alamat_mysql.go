package repository

import (
	"evermos/domain"

	"gorm.io/gorm"
)

type alamatMySQL struct {
	db *gorm.DB
}

func NewAlamatMySQL(db *gorm.DB) AlamatRepository {
	return &alamatMySQL{db}
}

func (r *alamatMySQL) Create(alamat *domain.Alamat) error {
	return r.db.Create(alamat).Error
}

func (r *alamatMySQL) FindByUserID(userID uint) ([]domain.Alamat, error) {
	var alamat []domain.Alamat
	err := r.db.Where("user_id = ?", userID).Find(&alamat).Error
	return alamat, err
}

func (r *alamatMySQL) FindByID(id uint) (*domain.Alamat, error) {
	var alamat domain.Alamat
	err := r.db.First(&alamat, id).Error
	return &alamat, err
}

func (r *alamatMySQL) Update(alamat *domain.Alamat) error {
	return r.db.Save(alamat).Error
}

func (r *alamatMySQL) Delete(id uint) error {
	return r.db.Delete(&domain.Alamat{}, id).Error
}
