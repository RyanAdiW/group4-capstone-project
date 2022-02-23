package request

import (
	"database/sql"
	"fmt"
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
	query := (`INSERT INTO requests (id_user, id_asset, id_status, request_date, return_date, description, created_at, updated_at) VALUES (?, ?, ?, now(), ?, ?, now(), now())`)

	statement, err := rr.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(request.Id_user, request.Id_asset, request.Id_status, request.Return_date, request.Description)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// get request by id
func (rr *requestRepo) GetById(id int) (entities.Request, error) {
	var request entities.Request
	if request.Return_date != "" {
		row := rr.db.QueryRow(`select r.id, r.id_user, r.id_asset, r.id_status, r.request_date, r.return_date, r.description, u.name as user_name, a.name as asset_name, c.description as category, a.avail_quantity, s.description as status 
		from requests r
		join users u on u.id = r.id_user
		join status_check s on s.id = r.id_status
		join assets a on a.id = r.id_asset
			join categories c on c.id = a.id_category
		where r.id = ? and r.deleted_at is nul`, id)
		err := row.Scan(&request.Id, &request.Id_user, &request.Id_asset, &request.Id_status, &request.Request_date, &request.Return_date, &request.Description, &request.User_name, &request.Asset_name, &request.Category, &request.Avail_quantity, &request.Status)
		if err != nil {
			log.Println(err)
			return request, err
		}
	} else {
		row := rr.db.QueryRow(`select r.id, r.id_user, r.id_asset, r.id_status, r.request_date, r.description, u.name as user_name, a.name as asset_name, c.description as category, a.avail_quantity, s.description as status 
		from requests r
		join users u on u.id = r.id_user
		join status_check s on s.id = r.id_status
		join assets a on a.id = r.id_asset
			join categories c on c.id = a.id_category
		where r.id = ?`, id)
		err := row.Scan(&request.Id, &request.Id_user, &request.Id_asset, &request.Id_status, &request.Request_date, &request.Description, &request.User_name, &request.Asset_name, &request.Category, &request.Avail_quantity, &request.Status)
		if err != nil {
			log.Println(err)
			return request, err
		}
	}

	return request, nil
}

// update request
func (rr *requestRepo) Update(request entities.Request, id int) error {
	query := `UPDATE requests SET`
	var bind []interface{}

	if request.Id_asset != 0 {
		bind = append(bind, request.Id_asset)
		query += " id_asset = ?,"
	}
	if request.Id_status != 0 {
		bind = append(bind, request.Id_status)
		query += " id_status = ?,"
	}

	if request.Return_date != "" {
		bind = append(bind, request.Return_date)
		query += " return_date = ?,"
	}
	if request.Description != "" {
		bind = append(bind, request.Description)
		query += " description = ?,"
	}

	bind = append(bind, id)
	query += " request_date = now(), updated_at = now() WHERE id = ? AND deleted_at is null"

	res, err := rr.db.Exec(query, bind...)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

// update available quantity asset
func (rr *requestRepo) UpdateAvailQty(quantity, id int) error {
	query := (`update requests r 
	join assets a on a.id = r.id_asset
	set avail_quantity = ?
	where r.id = ?`)

	res, err := rr.db.Exec(query, quantity, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

// get asset available qty based on id_asset
func (rr *requestRepo) GetAvailQty(id int) (entities.Request, error) {
	var request entities.Request

	row := rr.db.QueryRow(`select r.id, r.id_asset, a.initial_quantity, a.avail_quantity from requests r
	join assets a on a.id = r.id_asset
	where r.id = ? and r.deleted_at is null`, id)

	err := row.Scan(&request.Id, &request.Id_asset, &request.Initial_quantity, &request.Avail_quantity)
	if err != nil {
		return request, err
	}
	return request, nil
}

// get requests
func (rr *requestRepo) Get(request_date, status, filter_date string) ([]entities.Request, error) {
	var condition string
	var requests []entities.Request

	// var bind []interface{}

	if request_date == "latest" {
		condition += "order by r.request_date desc"
	} else if request_date == "oldest" {
		condition += "order by r.request_date asc"
	}

	res, err := rr.db.Query(`select r.id, r.id_user, r.id_asset, r.id_status, r.request_date, r.return_date, r.description, u.name as user_name, a.name as asset_name, c.description as category, a.avail_quantity, s.description as status 
	from requests r
	join users u on u.id = r.id_user
	join status_check s on s.id = r.id_status
	join assets a on a.id = r.id_asset
		join categories c on c.id = a.id_category
	where id_status != 8
	` + condition)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Close()
	for res.Next() {
		var request entities.Request

		err = res.Scan(&request.Id, &request.Id_user, &request.Id_asset, &request.Id_status, &request.Request_date, &request.Return_date, &request.Description, &request.User_name, &request.Asset_name, &request.Category, &request.Avail_quantity, &request.Status)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		requests = append(requests, request)
	}
	return requests, nil
}
