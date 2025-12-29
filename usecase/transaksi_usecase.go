package usecase

import (
	"errors"
	"evermos/domain"
	"evermos/repository"
)

type TransaksiUsecase struct {
	transaksiRepo repository.TransaksiRepository
	produkRepo    repository.ProdukRepository
	logRepo       repository.LogProdukRepository
}

func NewTransaksiUsecase(
	tr repository.TransaksiRepository,
	pr repository.ProdukRepository,
	lr repository.LogProdukRepository,
) *TransaksiUsecase {
	return &TransaksiUsecase{
		transaksiRepo: tr,
		produkRepo:    pr,
		logRepo:       lr,
	}
}

func (u *TransaksiUsecase) Create(
	userID, tokoID, alamatID uint,
	items []domain.TransaksiItem,
) error {

	if len(items) == 0 {
		return errors.New("items cannot be empty")
	}

	var total float64
	var logs []domain.LogProduk

	for i, item := range items {
		prod, err := u.produkRepo.FindByID(item.ProdukID)
		if err != nil {
			return err
		}

		if prod.Stock < item.Qty {
			return errors.New("stock not enough")
		}

		// kurangi stok
		prod.Stock -= item.Qty
		if err := u.produkRepo.Update(prod); err != nil {
			return err
		}

		// isi item transaksi
		items[i].NamaProduk = prod.Name
		items[i].HargaProduk = prod.Price
		items[i].Subtotal = prod.Price * float64(item.Qty)

		total += items[i].Subtotal

		// log produk (snapshot)
		logs = append(logs, domain.LogProduk{
			ProdukID: prod.ID,
			Nama:     prod.Name,
			Harga:    prod.Price,
			Qty:      item.Qty,
			Subtotal: items[i].Subtotal,
		})
	}

	tx := domain.Transaksi{
		UserID:   userID,
		TokoID:   tokoID,
		AlamatID: alamatID,
		Total:    total,
		Items:    items,
	}

	// create transaksi
	if err := u.transaksiRepo.Create(&tx); err != nil {
		return err
	}

	// simpan log produk
	for i := range logs {
		logs[i].TransaksiID = tx.ID
		if err := u.logRepo.Create(&logs[i]); err != nil {
			return err
		}
	}

	return nil
}

// ================= LIST =================

func (u *TransaksiUsecase) List(
	userID uint,
	page, limit int,
) ([]domain.Transaksi, int64, error) {
	return u.transaksiRepo.ListByUserPaginate(userID, page, limit)
}

// ================= FILTER =================

func (u *TransaksiUsecase) FilterByDate(
	userID uint,
	start, end string,
) ([]domain.Transaksi, error) {
	return u.transaksiRepo.FilterByDate(userID, start, end)
}
