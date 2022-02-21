package request

import (
	"sirclo/project/capstone/entities"
)

type RequestRepo interface {
	GetById(int) (entities.Request, error)
	Create(entities.Request) error
	Update(entities.Request, int) error
}
