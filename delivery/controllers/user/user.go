package user

import (
	"fmt"
	"net/http"
	"strconv"

	"sirclo/project/capstone/entities"

	response "sirclo/project/capstone/delivery/common"
	middlewares "sirclo/project/capstone/delivery/middleware"
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
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		password := []byte(userRequest.Password)

		hashedPassword, errEncrypt := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if errEncrypt != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to encrpyt password"))
		}
		existingEmail, errEmail := uc.repository.GetEmail()
		if errEmail != nil {
			return fmt.Errorf("error from repo")
		}

		user := entities.User{
			Name:         userRequest.Name,
			Email:        userRequest.Email,
			Password:     string(hashedPassword),
			Birth_date:   userRequest.Birth_date,
			Phone_number: userRequest.Phone_number,
			Photo:        userRequest.Photo,
			Gender:       userRequest.Gender,
			Address:      userRequest.Address,
		}

		for _, v := range existingEmail {
			if v.Email == user.Email {
				return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to join, email has been registered"))
			}
		}

		// create user to database
		err := uc.repository.Create(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create user"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create user"))
	}
}

// 2. get all user controller
func (uc UserController) GetUsersController() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.repository.Get()
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all user", users))
	}
}

// 3. get user by id
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

// 4. update user
func (uc UserController) UpdateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetEmail(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		userId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		tokenId, _ := middlewares.GetId(c)
		if userId != tokenId {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// binding data
		user := entities.User{}
		if errBind := c.Bind(&user); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		password := []byte(user.Password)

		hashedPassword, errEncrypt := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if errEncrypt != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to encrpyt password"))
		}
		user.Password = string(hashedPassword)
		// update user based on id to database
		errUpdate := uc.repository.Update(user, userId)
		if errUpdate != nil {
			fmt.Println(errUpdate)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update user"))
	}
}

// 5. delete user
func (uc UserController) DeleteUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetEmail(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		userId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		tokenId, _ := middlewares.GetId(c)
		if userId != tokenId {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// delete user based on id from database
		errDelete := uc.repository.Delete(userId)
		if errDelete != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "delete success"))
	}
}
