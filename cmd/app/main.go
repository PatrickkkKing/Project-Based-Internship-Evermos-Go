package main

import (
	"evermos/config"
	httpDelivery "evermos/delivery/http"
	"evermos/delivery/http/middleware"
	"evermos/domain"
	"evermos/repository"
	"evermos/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.InitDB()
	db.AutoMigrate(
		&domain.User{},
		&domain.Toko{},
		&domain.Alamat{},
		&domain.Kategori{},
		&domain.Produk{},
		&domain.Transaksi{},
		&domain.TransaksiItem{},
	)
	if err := db.Error; err != nil {
		panic(err)
	}

	app := fiber.New()

	// ================= REPOSITORY =================
	userRepo := repository.NewUserMySQL(db)
	tokoRepo := repository.NewTokoMySQL(db)
	alamatRepo := repository.NewAlamatMySQL(db)
	kategoriRepo := repository.NewKategoriMySQL(db)
	produkRepo := repository.NewProdukMySQL(db)
	transaksiRepo := repository.NewTransaksiMySQL(db)

	// ================= USECASE =================
	authUC := usecase.NewAuthUsecase(userRepo, tokoRepo)
	userUC := usecase.NewUserUsecase(userRepo)
	tokoUC := usecase.NewTokoUsecase(tokoRepo)
	alamatUC := usecase.NewAlamatUsecase(alamatRepo)
	kategoriUC := usecase.NewKategoriUsecase(kategoriRepo)
	produkUC := usecase.NewProdukUsecase(produkRepo)
	transaksiUC := usecase.NewTransaksiUsecase(transaksiRepo, produkRepo)

	// ================= HANDLER =================
	authHandler := httpDelivery.NewAuthHandler(authUC)
	userHandler := httpDelivery.NewUserHandler(userUC)
	tokoHandler := httpDelivery.NewTokoHandler(tokoUC)
	alamatHandler := httpDelivery.NewAlamatHandler(alamatUC)
	kategoriHandler := httpDelivery.NewKategoriHandler(kategoriUC)
	produkHandler := httpDelivery.NewProdukHandler(produkUC)
	transaksiHandler := httpDelivery.NewTransaksiHandler(transaksiUC)

	// ================= PUBLIC ROUTES =================
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	// ================= PROTECTED ROUTES =================
	api := app.Group("/api", middleware.JWTProtected(tokoRepo))

	api.Get("/users/me", userHandler.Me)
	api.Put("/users/me", userHandler.UpdateMe)

	api.Get("/toko/me", tokoHandler.MyToko)
	api.Put("/toko/me", tokoHandler.UpdateMyToko)

	api.Post("/alamat", alamatHandler.Create)
	api.Get("/alamat", alamatHandler.List)
	api.Put("/alamat/:id", alamatHandler.Update)
	api.Delete("/alamat/:id", alamatHandler.Delete)

	api.Get("/kategori", kategoriHandler.List)
	api.Post("/kategori", kategoriHandler.Create)
	api.Put("/kategori/:id", kategoriHandler.Update)
	api.Delete("/kategori/:id", kategoriHandler.Delete)

	app.Post("/register-admin", authHandler.RegisterAdmin)

	api.Get("/produk", produkHandler.List)
	api.Post("/produk", produkHandler.Create)
	api.Put("/produk/:id", produkHandler.Update)
	api.Delete("/produk/:id", produkHandler.Delete)

	api.Post("/transaksi", transaksiHandler.Create)
	api.Get("/transaksi", transaksiHandler.List)
	api.Get("/transaksi/filter", transaksiHandler.Filter)
	// db = db.Debug()
	// db.AutoMigrate(&domain.Alamat{})
	// db.AutoMigrate(&domain.Kategori{})
	// db.AutoMigrate(&domain.Produk{})
	// db.AutoMigrate(&domain.Transaksi{})
	// db.AutoMigrate(&domain.TransaksiItem{})

	app.Listen(":3000")
}
