package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sirclo/project/capstone/delivery/common"
	_middlewares "sirclo/project/capstone/delivery/middleware"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"password": "sasuke",
			"email":    "sasuke@mail.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := NewAuthController(mockAuthRepository{})
		if assert.NoError(t, (AuthController.LoginEmailController())(context)) {
			bodyResponses := res.Body.String()
			fmt.Println(bodyResponses)
			var response LoginResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Login because binding", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Password": "asdasd",
			"Email":    123,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := NewAuthController(mockAuthRepository{})
		if assert.NoError(t, (AuthController.LoginEmailController())(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Login because wrong email/password", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Password": "asdasd",
			"Email":    "sasukes@mail.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := NewAuthController(mockAuthRepository{})
		if assert.NoError(t, (AuthController.LoginEmailController())(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed create token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Password": "asdasd",
			"Email":    "sasuke",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := NewAuthController(mockAuthRepository{})
		if assert.NoError(t, (AuthController.LoginEmailController())(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
}

// =========================== mocking ===========================

type mockAuthRepository struct{}

func (m mockAuthRepository) LoginEmail(email, password string) (string, error) {
	if email == "sasuke@mail.com" {
		token, err := _middlewares.CreateToken(1, email, 1)
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", fmt.Errorf("failed create token")
}

func (m mockAuthRepository) GetPasswordByEmail(email string) (string, error) {
	if email == "sasuke@mail.com" {
		sandi, _ := bcrypt.GenerateFromPassword([]byte("sasuke"), bcrypt.DefaultCost)
		return string(sandi), nil
	}

	if email == "sasukes@mail.com" {
		sandi, _ := bcrypt.GenerateFromPassword([]byte("salah"), bcrypt.DefaultCost)
		return string(sandi), nil
	}

	return "", fmt.Errorf("no record")
}

func (m mockAuthRepository) GetIdByEmail(email string) (int, error) {
	if email == "sasuke@mail.com" {
		return 1, nil
	}

	return 0, fmt.Errorf("no record")
}

func (m mockAuthRepository) GetIdRole(email string) (int, error) {
	if email == "sasuke@mail.com" {
		return 1, nil
	}

	return 0, fmt.Errorf("no record")
}

func (m mockAuthRepository) GetNameByEmail(email string) (string, error) {
	if email == "sasuke@mail.com" {
		return "sasuke", nil
	}

	return "", fmt.Errorf("no record")
}
