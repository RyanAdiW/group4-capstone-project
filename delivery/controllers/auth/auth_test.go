package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sirclo/project/capstone/entities"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginEmailController(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Password": "asdasd",
			"Email":    "asd",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.Login
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

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.Login
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Login because wrong email/password", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Password": "asdasd",
			"Email":    "gagal",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.Login
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Login because failed hashing", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Password": "asdasd",
			"Email":    "gagalhash",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.Login
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Login because login function failed", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Password": "asdasd",
			"Email":    "gagallogin",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		AuthController := New(mockAuthRepository{})
		if assert.NoError(t, (AuthController.Login())(context)) {
			bodyResponses := res.Body.String()
			var response entities.Login
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
}

// =========================== mocking ===========================

type mockAuthRepository struct{}

func (m mockAuthRepository) Register(User entities.User) error {
	if User.Email == "gagalregister" {
		return fmt.Errorf("email not unique")
	}
	return nil
}

func (m mockAuthRepository) Login(email string) (entities.Login, error) {
	var gagal entities.Login
	if email == "gagallogin" {
		return gagal, fmt.Errorf("servernya mati")
	}
	// sukses
	return entities.Login{
		Id:    1,
		Name:  "asd",
		Email: "asd@gmail.com",
		Token: "asdasd",
	}, nil
}

func (m mockAuthRepository) FindUserByEmail(email string) (entities.User, error) {
	var gagal entities.User
	if email == "gagal" {
		return gagal, fmt.Errorf("user not found")
	}
	if email == "gagalhash" {
		return entities.User{
			Id:       1,
			Email:    "asd@gmail.com",
			Password: "asd",
		}, nil
	}
	sandi, _ := bcrypt.GenerateFromPassword([]byte("asdasd"), bcrypt.MinCost)
	return entities.User{
		Id:       1,
		Email:    "asd@gmail.com",
		Password: string(sandi),
	}, nil
}
