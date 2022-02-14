package user

import (
	"sirclo/project/capstone/entities"
)

type UserRepo interface {
	Create(entities.User) error
	Get() ([]entities.User, error)
	GetById(int) (entities.User, error)
	Update(entities.User, int) error
	Delete(int) error
}
