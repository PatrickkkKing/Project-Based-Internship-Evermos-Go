package usecase

import (
	"errors"
	"evermos/domain"
	"evermos/repository"
)

type AlamatUsecase struct {
	repo repository.AlamatRepository
}

func NewAlamatUsecase(repo repository.AlamatRepository) *AlamatUsecase {
	return &AlamatUsecase{repo}
}

func (u *AlamatUsecase) Create(userID uint, a *domain.Alamat) error {
	a.UserID = userID
	return u.repo.Create(a)
}

func (u *AlamatUsecase) List(userID uint) ([]domain.Alamat, error) {
	return u.repo.FindByUserID(userID)
}

func (u *AlamatUsecase) Update(userID uint, id uint, data *domain.Alamat) error {
	alamat, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	if alamat.UserID != userID {
		return errors.New("forbidden")
	}

	alamat.Label = data.Label
	alamat.Alamat = data.Alamat
	alamat.Kota = data.Kota
	alamat.Provinsi = data.Provinsi
	alamat.KodePos = data.KodePos

	return u.repo.Update(alamat)
}

func (u *AlamatUsecase) Delete(userID uint, id uint) error {
	alamat, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	if alamat.UserID != userID {
		return errors.New("forbidden")
	}

	return u.repo.Delete(id)
}
