package auth

import (
	"fmt"
	"net/http"
	"sirclo/project/capstone/delivery/common"
	"sirclo/project/capstone/repository/auth"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repository auth.Auth
}

func NewAuthController(repository auth.Auth) *AuthController {
	return &AuthController{repository: repository}
}

func (ac AuthController) LoginEmailController() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginRequest LoginEmailRequestFormat

		//bind request data
		if err := c.Bind(&loginRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "failed to bind"))
		}
		password := []byte(loginRequest.Password)

		hashedPassword, errPass := ac.repository.GetPasswordByEmail(loginRequest.Email)
		if errPass != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "email not found"))
		}

		errMatch := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
		if errMatch != nil {
			fmt.Println(errMatch)
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "failed to encrypt"))
		}

		// get token from login credential
		token, err := ac.repository.LoginEmail(loginRequest.Email, hashedPassword)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "failed to create token"))
		}

		uid, _ := ac.repository.GetIdByEmail(loginRequest.Email)
		idRole, _ := ac.repository.GetIdRole(loginRequest.Email)

		data := LoginResponseFormat{
			Token:         token,
			Id_user:       uid,
			Id_role:       idRole,
			Current_email: loginRequest.Email,
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("success", "login success", data))
	}
}
