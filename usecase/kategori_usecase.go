package usecase

import (
	"errors"
	"evermos/domain"
	"evermos/repository"
)

type KategoriUsecase struct {
	repo repository.KategoriRepository
}

func NewKategoriUsecase(repo repository.KategoriRepository) *KategoriUsecase {
	return &KategoriUsecase{repo}
}

func (u *KategoriUsecase) Create(role string, k *domain.Kategori) error {
	if role != "admin" {
		return errors.New("forbidden")
	}
	return u.repo.Create(k)
}

func (u *KategoriUsecase) List() ([]domain.Kategori, error) {
	return u.repo.List()
}

func (u *KategoriUsecase) Update(role string, id uint, k *domain.Kategori) error {
	if role != "admin" {
		return errors.New("forbidden")
	}
	kat, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	kat.Name = k.Name
	return u.repo.Update(kat)
}

func (u *KategoriUsecase) Delete(role string, id uint) error {
	if role != "admin" {
		return errors.New("forbidden")
	}
	return u.repo.Delete(id)
}
