package usecase

import (
	"errors"
	"evermos/domain"
	"evermos/repository"
)

type ProdukUsecase struct {
	repo repository.ProdukRepository
}

func NewProdukUsecase(repo repository.ProdukRepository) *ProdukUsecase {
	return &ProdukUsecase{repo}
}

func (u *ProdukUsecase) Create(tokoID uint, p *domain.Produk) error {
	p.TokoID = tokoID
	return u.repo.Create(p)
}

func (u *ProdukUsecase) List(tokoID uint) ([]domain.Produk, error) {
	return u.repo.List(tokoID)
}

func (u *ProdukUsecase) Update(tokoID uint, id uint, p *domain.Produk) error {
	prod, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	if prod.TokoID != tokoID {
		return errors.New("forbidden")
	}
	prod.Name = p.Name
	prod.Desc = p.Desc
	prod.Price = p.Price
	prod.Stock = p.Stock
	prod.Image = p.Image
	return u.repo.Update(prod)
}

func (u *ProdukUsecase) Delete(tokoID uint, id uint) error {
	prod, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	if prod.TokoID != tokoID {
		return errors.New("forbidden")
	}
	return u.repo.Delete(id)
}
