package usecases

import (
	"tamaragl/go-birthday-api/src/entities"
	"tamaragl/go-birthday-api/src/repositories"
	"time"
)

type GetUserUsecase struct {
	Repo repositories.Repository
}

type GetUserUsecaseInterface interface {
	GetUserBirthdayMessage(item string) (*entities.BirthdayMessage, error)
}

func NewGetUserUsecase(r repositories.Repository) *GetUserUsecase {
	return &GetUserUsecase{Repo: r}
}

func (u *GetUserUsecase) GetUserBirthdayMessage(item string) (*entities.BirthdayMessage, error) {
	user, err := u.Repo.GetItem(item)
	if err != nil {
		return nil, err
	}

	return entities.NewBirthdayMessage(user, time.Now())

}
