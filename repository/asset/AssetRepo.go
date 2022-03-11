package asset

import (
	"database/sql"
	"fmt"
	"log"

	"sirclo/project/capstone/entities"
)

type assetRepo struct {
	db *sql.DB
}

func NewAssetRepo(db *sql.DB) *assetRepo {
	return &assetRepo{db: db}
}

// create asset
func (ar *assetRepo) Create(asset entities.Asset) error {
	query := (`INSERT INTO assets (id_category, is_maintenance, name, description, initial_quantity, avail_quantity, photo, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, now(), now())`)

	statement, err := ar.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(asset.Id_category, asset.Is_maintenance, asset.Name, asset.Description, asset.Initial_quantity, asset.Avail_quantity, asset.Photo)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// get all asset with filter
func (ar *assetRepo) Get(category, maintenance, avail string, limit, offset int) ([]entities.Asset, error) {
	var condition string
	var condLimit string

	var bind []interface{}

	if category != "" {
		bind = append(bind, category)
		condition += " and c.id=? "
	}

	if avail != "" {
		switch avail {
		case "no":
			condition += " and a.avail_quantity = 0"
		case "yes":
			condition += " and a.avail_quantity > 0"
		}
	}

	if limit != 0 && offset == 0 {
		bind = append(bind, limit)
		condLimit += "limit ?"
	}

	if limit != 0 && offset != 0 {
		bind = append(bind, offset)
		bind = append(bind, limit)
		condLimit += "limit ?, ?"
	}

	var assets []entities.Asset
	results, err := ar.db.Query(`select a.id, a.id_category, a.is_maintenance, a.name, a.description, a.initial_quantity, a.avail_quantity, a.photo, c.description as category
								from assets a
								join categories c on c.id = a.id_category
								where a.deleted_at is null`+condition+` order by a.id asc `+condLimit, bind...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var asset entities.Asset

		err = results.Scan(&asset.Id, &asset.Id_category, &asset.Is_maintenance, &asset.Name, &asset.Description, &asset.Initial_quantity, &asset.Avail_quantity, &asset.Photo, &asset.Category)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		assets = append(assets, asset)
	}
	return assets, nil
}

// get asset by id
func (ar *assetRepo) GetById(id int) (entities.Asset, error) {
	var asset entities.Asset

	row := ar.db.QueryRow(`select a.id, a.id_category, a.is_maintenance, a.name, a.description, a.initial_quantity, a.avail_quantity, a.photo, c.description as category
							from assets a
							join categories c on c.id = a.id_category
							where a.deleted_at is null and a.id = ? order by a.id asc`, id)

	err := row.Scan(&asset.Id, &asset.Id_category, &asset.Is_maintenance, &asset.Name, &asset.Description, &asset.Initial_quantity, &asset.Avail_quantity, &asset.Photo, &asset.Category)
	if err != nil {
		return asset, err
	}

	return asset, nil
}

// update asset
func (ar *assetRepo) Update(assetExisted, asset entities.Asset, id int) error {
	query := `UPDATE assets SET`
	var bind []interface{}

	if asset.Name != "" {
		bind = append(bind, asset.Name)
		query += " name = ?,"
	}

	if asset.Description != "" {
		bind = append(bind, asset.Description)
		query += " description = ?,"
	}

	if asset.Id_category != 0 {
		bind = append(bind, asset.Id_category)
		query += " id_category = ?,"
	}

	if asset.Initial_quantity != 0 {
		bind = append(bind, asset.Initial_quantity)
		query += " initial_quantity = ?,"

		if asset.Initial_quantity > assetExisted.Initial_quantity {
			diff := asset.Initial_quantity - assetExisted.Initial_quantity
			updAvail := diff + assetExisted.Avail_quantity
			bind = append(bind, updAvail)
			query += " avail_quantity = ?,"
			assetExisted.Avail_quantity = updAvail
		}

		if asset.Initial_quantity < assetExisted.Avail_quantity {
			return fmt.Errorf("Your quantity is lower than expected!")
		}
	}

	if asset.Photo != "" {
		bind = append(bind, asset.Photo)
		query += " photo = ?,"
	}

	if asset.Is_maintenance == true {
		if asset.Initial_quantity == 0 {
			if assetExisted.Avail_quantity == assetExisted.Initial_quantity {
				bind = append(bind, asset.Is_maintenance)
				query += " is_maintenance = ?,"

				bind = append(bind, 0)
				query += " avail_quantity = ?,"
			} else {
				return fmt.Errorf("All asset must be in maintenence!")
			}
		} else {
			if assetExisted.Avail_quantity == asset.Initial_quantity {
				bind = append(bind, asset.Is_maintenance)
				query += " is_maintenance = ?,"

				bind = append(bind, 0)
				query += " avail_quantity = ?,"
			} else {
				return fmt.Errorf("All asset must be in maintenence!")
			}
		}
	} else {
		if asset.Initial_quantity == 0 {
			bind = append(bind, asset.Is_maintenance)
			query += " is_maintenance = ?,"

			bind = append(bind, assetExisted.Initial_quantity)
			query += " avail_quantity = ?,"
		} else {
			bind = append(bind, asset.Is_maintenance)
			query += " is_maintenance = ?,"

			bind = append(bind, asset.Initial_quantity)
			query += " avail_quantity = ?,"
		}
	}

	bind = append(bind, id)
	query += " updated_at = now() WHERE id = ? AND deleted_at is null"

	res, err := ar.db.Exec(query, bind...)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

// delete asset
func (ar *assetRepo) Delete(id int) error {
	res, err := ar.db.Exec("UPDATE assets SET deleted_at = now() WHERE id = ? AND deleted_at is null", id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

func (ar *assetRepo) GetSummaryAsset() (entities.SummaryAsset, error) {
	var summary entities.SummaryAsset
	row := ar.db.QueryRow(`select all_asset.total_asset, all_asset.total_avail_asset, COALESCE(maintenance.total_asset_maintenance, 0)
							from (
								SELECT 
									sum(a.initial_quantity) as total_asset, sum(a.avail_quantity) as total_avail_asset
								FROM
									assets a
									join categories c on c.id = a.id_category
									where a.deleted_at is null order by a.id asc) AS all_asset
							JOIN (
								SELECT 
									sum(a.initial_quantity) as total_asset_maintenance
								FROM
									assets a
									join categories c on c.id = a.id_category
									where a.deleted_at is null and a.is_maintenance = true order by a.id asc
							) AS maintenance
							`)

	err := row.Scan(&summary.Total_asset, &summary.Available, &summary.Maintenance)
	if err != nil {
		return summary, err
	}
	summary.Use = summary.Total_asset - summary.Available - summary.Maintenance

	return summary, nil
}

func (ar *assetRepo) GetHistoryUsage(id_asset, limit, offset int) (entities.HistoryUsage, error) {
	var condLimit string
	var bind []interface{}

	bind = append(bind, id_asset)
	if limit != 0 && offset == 0 {
		bind = append(bind, limit)
		condLimit += "limit ?"
	}

	if limit != 0 && offset != 0 {
		bind = append(bind, offset)
		bind = append(bind, limit)
		condLimit += "limit ?, ?"
	}

	var historyUsage entities.HistoryUsage
	results, err := ar.db.Query(`select a.id, a.id_category, a.name as asset_name, a.description, a.photo, c.description as category,
									r.id, r.id_user, r.id_asset, u.name as user_name, r.request_date, s.description as status
								from assets a
								join categories c on c.id = a.id_category and c.deleted_at is null
								left join requests r on r.id_asset = a.id and r.deleted_at is null
								join users u on u.id = r.id_user
								join status_check s on s.id = r.id_status
								WHERE a.id = ? `+condLimit, bind...)
	if err != nil {
		log.Println(err)
		return historyUsage, err
	}

	defer results.Close()

	for results.Next() {
		var history entities.History

		err = results.Scan(&historyUsage.Id, &historyUsage.Id_category, &historyUsage.Name, &historyUsage.Description,
			&historyUsage.Photo, &historyUsage.Category, &history.Id, &history.Id_user,
			&history.Id_asset, &history.Name, &history.Request_date, &history.Status)
		if err != nil {
			log.Println(err)
			return historyUsage, err
		}

		historyUsage.List_history = append(historyUsage.List_history, history)
	}

	return historyUsage, nil
}

func (ar *assetRepo) GetCategory() ([]entities.Categories, error) {
	var categories []entities.Categories
	row, err := ar.db.Query(`select id, description
						from categories
						where deleted_at is null`)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		var category entities.Categories

		err = row.Scan(&category.Id, &category.Description)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
