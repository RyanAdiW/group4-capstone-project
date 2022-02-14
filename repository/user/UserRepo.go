package user

import (
	"database/sql"
	"fmt"

	"sirclo/project/capstone/entities"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

// create user
func (ur *userRepo) Create(user entities.User) error {
	query := (`INSERT INTO users (name, email, password, birth_date, phone_number, photo, gender, address, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, now(), now())`)

	statement, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password, user.Birth_date, user.Phone_number, user.Photo, user.Gender, user.Address)
	if err != nil {
		return err
	}

	return nil
}

// get all user
func (ur *userRepo) Get() ([]entities.User, error) {
	var users []entities.User
	results, err := ur.db.Query("select id, name, email, birth_date, phone_number, photo, gender, address from users where deleted_at is null order by id asc")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var user entities.User

		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Birth_date, &user.Phone_number, &user.Photo, &user.Gender, &user.Address)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

// get user by id
func (ur *userRepo) GetById(id int) (entities.User, error) {
	var user entities.User

	row := ur.db.QueryRow(`SELECT id, name, email, birth_date, phone_number, photo, gender, address FROM users WHERE id = ? AND deleted_at IS NULL`, id)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Birth_date, &user.Phone_number, &user.Photo, &user.Gender, &user.Address)
	if err != nil {
		return user, err
	}

	return user, nil
}

// update user
func (ur *userRepo) Update(user entities.User, id int) error {
	res, err := ur.db.Exec("UPDATE users SET name = ?, email = ?, birth_date = ?, phone_number = ?, photo = ?, gender = ?, address = ?, updated_at = now() WHERE id = ?", user.Name, user.Email, user.Birth_date, user.Phone_number, user.Photo, user.Gender, user.Address, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

// delete user
func (ur *userRepo) Delete(id int) error {
	res, err := ur.db.Exec("UPDATE users SET deleted_at = now() WHERE id = ? AND deleted_at is null", id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}
