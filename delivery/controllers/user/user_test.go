package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	middlewares "sirclo/project/capstone/delivery/middleware"
	"sirclo/project/capstone/entities"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// 1. test create user
func TestCreateUser(t *testing.T) {
	t.Run("failed to bind data", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "asd",
			"email":    "asd@mail.com",
			"password": 12345,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, userController.CreateUserController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to bind data", response.Message)
		}
	})
	t.Run("email has been registered", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "asd",
			"email":    "asd@mail.com",
			"password": "12345",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, userController.CreateUserController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to join, email has been registered", response.Message)
		}
	})
	t.Run("failed to create user", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "qwer",
			"email":    "qwer@mail.com",
			"password": "12345",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, userController.CreateUserController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to create user", response.Message)
		}
	})
	t.Run("succes create user", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "qwer",
			"email":    "qwer@mail.com",
			"password": "12345",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, userController.CreateUserController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success create user", response.Message)
		}
	})
}

// 2. Get user by id
func TestGetUser(t *testing.T) {
	var (
		globalToken, errCreateToken = middlewares.CreateToken(2, "admin", 1)
	)
	t.Run("success get user by id", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetByIdController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get user", response.Message)
		}

	})
	t.Run("failed to convert id", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("a")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetByIdController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert id", response.Message)
		}
	})
	t.Run("failed to fetch data", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := NewUserController(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetByIdController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to fetch data", response.Message)
		}

	})
}

// 3. test get all user
func TestGetAllUser(t *testing.T) {
	var (
		globalToken, errCreateToken = middlewares.CreateToken(2, "admin", 1)
	)
	t.Run("success get all user", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUsersController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all user", response.Message)
		}

	})
	t.Run("failed to fetch data", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUsersController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to fetch data", response.Message)
		}

	})
}

type mockUserRepository struct{}

func (m mockUserRepository) Create(entities.User) error {
	return nil
}
func (m mockUserRepository) GetById(id int) (entities.User, error) {
	return entities.User{}, nil
}
func (m mockUserRepository) Get() ([]entities.User, error) {
	return []entities.User{
		{Id: 1,
			Name:  "asd",
			Email: "asd@mail.com",
		},
		{Id: 2,
			Name:  "dsa",
			Email: "dsa@mail.com",
		},
	}, nil
}
func (m mockUserRepository) GetEmail() ([]entities.User, error) {
	return []entities.User{
		{Id: 1,
			Email: "asd@mail.com",
		},
		{Id: 2,
			Email: "dsa@mail.com",
		},
	}, nil
}

type mockErrorUserRepository struct{}

func (m mockErrorUserRepository) Create(entities.User) error {
	return fmt.Errorf("error")
}
func (m mockErrorUserRepository) GetById(id int) (entities.User, error) {
	return entities.User{}, fmt.Errorf("error")
}
func (m mockErrorUserRepository) Get() ([]entities.User, error) {
	return []entities.User{
		{Id: 1,
			Name:  "asd",
			Email: "asd@mail.com",
		},
		{Id: 2,
			Name:  "dsa",
			Email: "dsa@mail.com",
		},
	}, fmt.Errorf("error")
}
func (m mockErrorUserRepository) GetEmail() ([]entities.User, error) {
	return []entities.User{
		{Id: 1,
			Email: "asd@mail.com",
		},
		{Id: 2,
			Email: "dsa@mail.com",
		},
	}, nil
}
