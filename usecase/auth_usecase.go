package usecase

import (
	"errors"
	"evermos/config"
	"evermos/domain"
	"evermos/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo repository.UserRepository
	tokoRepo repository.TokoRepository
}

func NewAuthUsecase(userRepo repository.UserRepository, tokoRepo repository.TokoRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
		tokoRepo: tokoRepo,
	}
}

// ================= REGISTER =================
func (u *AuthUsecase) Register(name, email, phone, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := domain.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: string(hash),
		Role:     "user",
	}

	if err := u.userRepo.Create(&user); err != nil {
		return err
	}

	toko := domain.Toko{
		UserID: user.ID,
		Name:   name + " Store",
	}

	return u.tokoRepo.Create(&toko)
}

func (u *AuthUsecase) RegisterAdmin(name, email, phone, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := domain.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: string(hash),
		Role:     "admin",
		// Toko: domain.Toko{
		// 	Name: name + " Store",
		// },
	}

	return u.userRepo.Create(&user)
}

// ================= LOGIN =================
func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email or password wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email or password wrong")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
