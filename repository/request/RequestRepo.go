package request

import (
	"database/sql"
	"sirclo/project/capstone/entities"
)

type requestRepo struct {
	db *sql.DB
}

func NewRequestRepo(db *sql.DB) *requestRepo {
	return &requestRepo{db: db}
}

// create request
func (rr *requestRepo) Create(request entities.Request) error {
	query := (`INSERT INTO requests (id_user, id_asset, id_status, request_date, return_date,)`)
}
