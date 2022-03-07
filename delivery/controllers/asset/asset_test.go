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

// 1. test create asset
func TestCreateAsset(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 3)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             "laptop",
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets")

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
			"name":             1,
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets")

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
	t.Run("failed to create asset", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             "laptop",
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      10,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets")

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
			assert.Equal(t, "failed to create asset", response.Message)
		}
	})
	t.Run("success to create asset", func(t *testing.T) {
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
		context.SetPath("/assets")

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
			assert.Equal(t, "success create asset", response.Message)
		}
	})
}

// 2. test get assets
func TestGetAssets(t *testing.T) {
	t.Run("failed to fetch data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/?category=100", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetAssetsController())(context)) {
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
	t.Run("success get assets", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/?limit=1", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetAssetsController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all assets", response.Message)
		}
	})
}

// 3. test get asset by id
func TestGetAssetById(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/detail")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetAssetByIdController())(context)) {
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
	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/detail")
		context.SetParamNames("id")
		context.SetParamValues("a")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetAssetByIdController())(context)) {
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
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/detail")
		context.SetParamNames("id")
		context.SetParamValues("100")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetAssetByIdController())(context)) {
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
	t.Run("success get asset", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/detail")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetAssetByIdController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get asset", response.Message)
		}
	})
}

// 4. test update asset
func TestUpdateAsset(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             "laptop",
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/update")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateAssetController())(context)) {
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

	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             "laptop",
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/update")
		context.SetParamNames("id")
		context.SetParamValues("a")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateAssetController())(context)) {
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
	t.Run("failed to bind data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             12,
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/update")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateAssetController())(context)) {
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

	t.Run("failed get data assets", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             "laptop",
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/update")
		context.SetParamNames("id")
		context.SetParamValues("100")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusInternalServerError, res.Code)
			assert.Equal(t, "error", response.Status)
			assert.Equal(t, "err get data asset", response.Message)
		}
	})

	t.Run("failed get data assets", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             "",
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      100,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/update")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed update data", response.Message)
		}
	})

	t.Run("success update request", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":             "laptop",
			"description":      "macbook m1 pro",
			"initial_quantity": 1,
			"photo":            "",
			"id_category":      1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/update")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update asset", response.Message)
		}
	})

}

// 5. test delete asset
func TestDeleteAsset(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/delete")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.DeleteAssetController())(context)) {
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

	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/delete")
		context.SetParamNames("id")
		context.SetParamValues("a")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.DeleteAssetController())(context)) {
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

	t.Run("failed to delete data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/delete")
		context.SetParamNames("id")
		context.SetParamValues("100")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.DeleteAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "data not found", response.Message)
		}
	})

	t.Run("success delete asset", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/delete")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.DeleteAssetController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "delete success", response.Message)
		}
	})

}

// 6. test get summary asset
func TestGetSummaryAsset(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/summary")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetSummaryAssetsController())(context)) {
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

	t.Run("failed to fetch data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/summary")

		reqController := NewAssetController(mockErrorAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetSummaryAssetsController())(context)) {
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

	t.Run("success get summary asset", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/summary")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetSummaryAssetsController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get summary assets", response.Message)
		}
	})

}

// 7. test get history usage asset
func TestGetHistoryUsage(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/usage")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetHistoryUsageController())(context)) {
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
	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/usage")
		context.SetParamNames("id")
		context.SetParamValues("a")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetHistoryUsageController())(context)) {
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
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/usage")
		context.SetParamNames("id")
		context.SetParamValues("100")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetHistoryUsageController())(context)) {
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
	t.Run("success get asset", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/usage")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetHistoryUsageController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get history", response.Message)
		}
	})
}

// 8. test get all categories asset
func TestGetCategories(t *testing.T) {

	t.Run("failed to fetch data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/categories")

		reqController := NewAssetController(mockErrorAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetCategoriesController())(context)) {
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

	t.Run("success get summary asset", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/assets/categories")

		reqController := NewAssetController(mockAssetRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetCategoriesController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get categories", response.Message)
		}
	})

}

type mockAssetRepository struct{}

func (m mockAssetRepository) Create(asset entities.Asset) error {
	if asset.Id_category == 10 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockAssetRepository) Get(category, maintenance, avail string, limit, offset int) ([]entities.Asset, error) {
	if category == "100" {
		return nil, fmt.Errorf("error")
	}
	return nil, nil
}

func (m mockAssetRepository) GetById(id int) (entities.Asset, error) {
	if id == 100 {
		return entities.Asset{}, fmt.Errorf("error")
	}
	return entities.Asset{}, nil
}

func (m mockAssetRepository) Update(assetExisted, asset entities.Asset, id int) error {
	if asset.Id_category == 100 {
		return fmt.Errorf("errors")
	}
	return nil
}

func (m mockAssetRepository) Delete(id int) error {
	if id == 100 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockAssetRepository) GetSummaryAsset() (entities.SummaryAsset, error) {
	return entities.SummaryAsset{}, nil
}

func (m mockAssetRepository) GetHistoryUsage(id_asset, limit, offset int) (entities.HistoryUsage, error) {
	if id_asset == 100 {
		return entities.HistoryUsage{}, fmt.Errorf("error")
	}
	return entities.HistoryUsage{}, nil
}

func (m mockAssetRepository) GetCategory() ([]entities.Categories, error) {
	return nil, nil
}

type mockErrorAssetRepository struct{}

func (m mockErrorAssetRepository) Create(asset entities.Asset) error {
	if asset.Id_category == 10 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockErrorAssetRepository) Get(category, maintenance, avail string, limit, offset int) ([]entities.Asset, error) {
	if category == "100" {
		return nil, fmt.Errorf("error")
	}
	return nil, nil
}

func (m mockErrorAssetRepository) GetById(id int) (entities.Asset, error) {
	if id == 100 {
		return entities.Asset{}, fmt.Errorf("error")
	}
	return entities.Asset{}, nil
}

func (m mockErrorAssetRepository) Update(assetExisted, asset entities.Asset, id int) error {
	if asset.Id_category == 100 {
		return fmt.Errorf("errors")
	}
	return nil
}

func (m mockErrorAssetRepository) Delete(id int) error {
	if id == 100 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockErrorAssetRepository) GetSummaryAsset() (entities.SummaryAsset, error) {
	return entities.SummaryAsset{}, fmt.Errorf("error")
}

func (m mockErrorAssetRepository) GetHistoryUsage(int, int, int) (entities.HistoryUsage, error) {
	return entities.HistoryUsage{}, nil
}

func (m mockErrorAssetRepository) GetCategory() ([]entities.Categories, error) {
	return nil, fmt.Errorf("error")
}
