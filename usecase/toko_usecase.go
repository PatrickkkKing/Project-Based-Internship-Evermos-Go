package usecase

import (
	"errors"
	"evermos/domain"
	"evermos/repository"
)

type TokoUsecase struct {
	tokoRepo repository.TokoRepository
}

func NewTokoUsecase(tokoRepo repository.TokoRepository) *TokoUsecase {
	return &TokoUsecase{tokoRepo}
}

func (u *TokoUsecase) GetMyToko(userID uint) (*domain.Toko, error) {
	return u.tokoRepo.FindByUserID(userID)
}

func (u *TokoUsecase) UpdateMyToko(userID uint, name string) error {
	toko, err := u.tokoRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("toko not found")
	}

	toko.Name = name
	return u.tokoRepo.Update(toko)
}
