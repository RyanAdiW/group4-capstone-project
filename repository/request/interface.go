package request

import (
	"sirclo/project/capstone/entities"
)

type RequestRepo interface {
	Create(entities.Request) error
	Get() ([]entities.Request, error)
	GetById(int) (entities.Request, error)
	Update(entities.Request, int) error
	UpdateAvailQty(int, int) error
	GetAvailQty(int) (entities.Request, error)
}
