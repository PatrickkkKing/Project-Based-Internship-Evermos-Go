package usecase

import (
	"errors"
	"evermos/domain"
	"evermos/repository"
)

type TransaksiUsecase struct {
	transaksiRepo repository.TransaksiRepository
	produkRepo    repository.ProdukRepository
}

func NewTransaksiUsecase(
	tr repository.TransaksiRepository,
	pr repository.ProdukRepository,
) *TransaksiUsecase {
	return &TransaksiUsecase{tr, pr}
}

func (u *TransaksiUsecase) Create(
	userID, tokoID, alamatID uint,
	items []domain.TransaksiItem,
) error {

	var total float64

	for i, item := range items {
		prod, err := u.produkRepo.FindByID(item.ProdukID)
		if err != nil {
			return err
		}

		if prod.Stock < item.Qty {
			return errors.New("stock not enough")
		}

		prod.Stock -= item.Qty
		u.produkRepo.Update(prod)

		items[i].NamaProduk = prod.Name
		items[i].HargaProduk = prod.Price
		items[i].Subtotal = prod.Price * float64(item.Qty)

		total += items[i].Subtotal
	}

	tx := domain.Transaksi{
		UserID:   userID,
		TokoID:   tokoID,
		AlamatID: alamatID,
		Total:    total,
		Items:    items,
	}

	return u.transaksiRepo.Create(&tx)
}

func (u *TransaksiUsecase) ListByUser(userID uint) ([]domain.Transaksi, error) {
	return u.transaksiRepo.ListByUser(userID)
}

func (u *TransaksiUsecase) List(
	userID uint,
	page, limit int,
) ([]domain.Transaksi, int64, error) {
	return u.transaksiRepo.ListByUserPaginate(userID, page, limit)
}

func (u *TransaksiUsecase) FilterByDate(
	userID uint,
	start, end string,
) ([]domain.Transaksi, error) {
	return u.transaksiRepo.FilterByDate(userID, start, end)
}
