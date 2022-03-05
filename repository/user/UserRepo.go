package user

import (
	"database/sql"
	"fmt"
	"log"

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
	query := (`INSERT INTO users (name, email, password, divisi, created_at, updated_at, id_role) VALUES (?, ?, ?, ?, now(), now(), ?)`)

	statement, err := ur.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password, user.Divisi, user.Id_role)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// get all user
func (ur *userRepo) Get() ([]entities.User, error) {
	var users []entities.User
	results, err := ur.db.Query(`select u.id, u.name, u.email, u.divisi, r.id as id_role, r.description as role
								from users u
								join roles r on r.id=u.id_role
								where u.deleted_at is null and u.id_role = 2 order by u.id asc`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var user entities.User

		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Divisi, &user.Id_role, &user.Role)
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

	row := ur.db.QueryRow(`select u.id, u.name, u.email, u.divisi, r.id as id_role, r.description as role
							from users u
							join roles r on r.id=u.id_role
							WHERE u.id = ? AND u.deleted_at IS NULL`, id)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Divisi, &user.Id_role, &user.Role)
	if err != nil {
		return user, err
	}

	return user, nil
}

// get email user
func (ur *userRepo) GetEmail() ([]entities.User, error) {
	var users []entities.User
	res, err := ur.db.Query("select email from users where deleted_at is null")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var user entities.User

		err = res.Scan(&user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
