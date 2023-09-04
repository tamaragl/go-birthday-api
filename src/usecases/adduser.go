package usecases

import (
	"tamaragl/go-birthday-api/src/entities"
	"tamaragl/go-birthday-api/src/repositories"
)

type AddUserUsecase struct {
	Repo repositories.Repository
}

type AddUserUsecaseInterface interface {
	Add(user *entities.User) error
}

func NewAddUserUsecase(r repositories.Repository) *AddUserUsecase {
	return &AddUserUsecase{Repo: r}
}

func (u *AddUserUsecase) Add(user *entities.User) error {
	return u.Repo.PutItem(user)
}
