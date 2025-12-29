package usecase

import (
	"errors"
	"evermos/domain"
	"evermos/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo}
}

func (u *UserUsecase) GetMe(userID uint) (*domain.User, error) {
	return u.userRepo.FindByID(userID)
}

func (u *UserUsecase) UpdateMe(userID uint, name, phone string) error {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if user.ID != userID {
		return errors.New("forbidden")
	}

	user.Name = name
	user.Phone = phone

	return u.userRepo.Update(user)
}
