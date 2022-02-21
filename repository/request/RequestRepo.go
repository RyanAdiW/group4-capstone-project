package request

import (
	"database/sql"
	"log"
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
	query := (`INSERT INTO requests (id_user, id_asset, id_status, request_date, return_date, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, now(), now())`)

	statement, err := rr.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(request.Id_user, request.Id_asset, request.Id_status, request.Request_date, request.Return_date, request.Description)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// func (rr *requestRepo) GetById(id int) (entities.Request, error) {

// }
