package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"sirclo/project/capstone/entities"

	response "sirclo/project/capstone/delivery/common"
	userRepo "sirclo/project/capstone/repository/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repository userRepo.UserRepo
}

func NewUserController(user userRepo.UserRepo) *UserController {
	return &UserController{repository: user}
}

// 1. create user controller
func (uc UserController) CreateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// bind data
		var userRequest UserRequestFormat
		if err := c.Bind(&userRequest); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		//set password
		password := []byte(userRequest.Password)

		hashedPassword, errEncrypt := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if errEncrypt != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to encrpyt password"))
		}
		existingEmail, errEmail := uc.repository.GetEmail()
		if errEmail != nil {
			log.Println("error from repo")
			return fmt.Errorf("error from repo")
		}

		user := entities.User{
			Name:     userRequest.Name,
			Email:    userRequest.Email,
			Password: string(hashedPassword),
			Divisi:   userRequest.Divisi,
			Id_role:  userRequest.Id_role,
		}

		for _, v := range existingEmail {
			if v.Email == user.Email {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to join, email has been registered"))
			}
		}

		// create user to database
		err := uc.repository.Create(user)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create user"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create user"))
	}
}

// 2. get user by id
func (uc UserController) GetByIdController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get id from param
		userId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// get user from db
		user, err := uc.repository.GetById(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get user", user))
	}
}
