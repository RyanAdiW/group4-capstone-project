package asset

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	middlewares "sirclo/project/capstone/delivery/middleware"
	"sirclo/project/capstone/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// 1. test create request
func TestCreateAsset(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 3)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_user":     1,
			"id_asset":    1,
			"id_status":   1,
			"return_date": "2022-02-14",
			"description": "create",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "unauthorized access", response.Message)
		}
	})
	t.Run("failed to bind", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_user":     "1",
			"id_asset":    1,
			"id_status":   1,
			"return_date": "2022-02-14",
			"description": "create",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateAssetController())(context)) {
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
	t.Run("failed to create request", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_user":     1,
			"id_asset":    1,
			"id_status":   1,
			"return_date": "2022-02-14",
			"description": "create",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to create request", response.Message)
		}
	})
	t.Run("success to create request role 1", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_user":     1,
			"id_asset":    1,
			"id_status":   1,
			"return_date": "2022-02-14",
			"description": "create",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success create request", response.Message)
		}
	})
	t.Run("success to create request role 2", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 2)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_asset":    1,
			"id_status":   1,
			"return_date": "2022-02-14",
			"description": "create",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success create request", response.Message)
		}
	})
}

type mockAssetRepository struct{}

func (m mockAssetRepository) Create(entities.Asset) error {
	return nil
}

func (m mockAssetRepository) Get(string, string, string, int, int) ([]entities.Asset, error) {
	return nil, nil
}

func (m mockAssetRepository) GetById(int) (entities.Asset, error) {
	return entities.Asset{}, nil
}

func (m mockAssetRepository) Update(entities.Asset, entities.Asset, int) error {
	return nil
}

func (m mockAssetRepository) Delete(int) error {
	return nil
}

func (m mockAssetRepository) GetSummaryAsset() (entities.SummaryAsset, error) {
	return entities.SummaryAsset{}, nil
}

func (m mockAssetRepository) GetHistoryUsage(int, int, int) (entities.HistoryUsage, error) {
	return entities.HistoryUsage{}, nil
}

func (m mockAssetRepository) GetCategory() ([]entities.Categories, error) {
	return nil, nil
}
