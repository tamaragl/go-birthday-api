package repositories

import "tamaragl/go-birthday-api/src/entities"

type Repository interface {
	GetItem(item string) (*entities.User, error)
	PutItem(user *entities.User) error
}
