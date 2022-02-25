package request

import (
	"sirclo/project/capstone/entities"
)

type RequestRepo interface {
	Create(entities.Request) error
	GetAdmin(request_date, status, filter_date string) ([]entities.RequestResponse, error)
	GetManager(request_date, status, filter_date string) ([]entities.RequestResponse, error)
	GetById(int) (entities.RequestResponse, error)
	Update(entities.Request, int) error
	UpdateAvailQty(int, int) error
	GetAvailQty(int) (entities.Request, error)
}
