package request

import (
	"sirclo/project/capstone/entities"
)

type RequestRepo interface {
	Create(entities.Request) error
	GetAdmin(returnDate, requestDate, status, filterDate, category string, limit, offset int) ([]entities.RequestResponse, error)
	GetManager(returnDate, requestDate, status, filterDate, category string, limit, offset int) ([]entities.RequestResponse, error)
	GetById(int) (entities.RequestResponse, error)
	Update(entities.Request, int) error
	UpdateAvailQty(int, int) error
	GetAvailQty(int) (entities.Request, error)
	GetEmployee(idEmployee int, isHistory bool, limit, offset int) ([]entities.RequestResponse, error)
}
