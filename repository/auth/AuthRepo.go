package auth

import (
	"database/sql"
	"fmt"

	"sirclo/project/capstone/entities"

	_middlewares "sirclo/project/capstone/delivery/middleware"
)

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *authRepo {
	return &authRepo{db: db}
}

func (ar *authRepo) LoginEmail(email, password string) (string, error) {
	result, err := ar.db.Query("select id, name, email, divisi, id_role FROM users WHERE email=? AND password=?", email, password)
	if err != nil {
		return "", err
	}
	if isExist := result.Next(); !isExist {
		return "", fmt.Errorf("id not found")
	}
	var user entities.User
	errScan := result.Scan(&user.Id, &user.Name, &user.Email, &user.Divisi, &user.Id_role)
	if errScan != nil {
		return "", errScan
	}
	token, err := _middlewares.CreateToken(user.Id, user.Email, user.Id_role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ar *authRepo) GetPasswordByEmail(email string) (string, error) {
	result, err := ar.db.Query("select password FROM users WHERE email=?", email)
	if err != nil {
		return "", err
	}
	if isExist := result.Next(); !isExist {
		return "", fmt.Errorf("id not found")
	}
	var user entities.User
	errScan := result.Scan(&user.Password)
	if errScan != nil {
		return "", errScan
	}
	password := user.Password
	return password, nil
}

func (ar *authRepo) GetIdByEmail(email string) (int, error) {
	result, err := ar.db.Query("SELECT id FROM users WHERE email=?", email)
	if err != nil {
		return 0, err
	}
	if isExist := result.Next(); !isExist {
		return 0, fmt.Errorf("id not found")
	}
	var user entities.User
	errScan := result.Scan(&user.Id)
	if errScan != nil {
		return 0, errScan
	}
	userId := user.Id
	return userId, nil
}
