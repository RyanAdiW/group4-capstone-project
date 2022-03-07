package request

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
func TestCreateRequest(t *testing.T) {
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

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateRequestEmployee())(context)) {
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

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateRequestEmployee())(context)) {
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

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateRequestEmployee())(context)) {
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

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateRequestEmployee())(context)) {
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

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.CreateRequestEmployee())(context)) {
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

// 2. test get request detail by id
func TestGetRequestById(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestByIdController())(context)) {
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
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("a")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestByIdController())(context)) {
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
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestByIdController())(context)) {
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
	t.Run("success get request", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestByIdController())(context)) {
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

// 3. test update status request
func TestUpdateRequest(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_status": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
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
	t.Run("unauthorized admin", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_status": 8,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "id_status must be 2 || 5 || 6 || 7", response.Message)
		}
	})
	t.Run("unauthorized employee", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 2)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_status": 7,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "id_status must be 8", response.Message)
		}
	})
	t.Run("unauthorized manager", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 3)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_status": 5,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "id_status must be 3 || 4", response.Message)
		}
	})
	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_status": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("a")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
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
			"id_status": "1",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
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
	t.Run("id status = 6", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_status": 6,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("success update request", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"id_status": 2,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.UpdateRequestStatus())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success update request", response.Message)
		}
	})

}

// 4. test get requests
func TestGetRequests(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 0)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")
		context.SetParamNames("id")
		context.SetParamValues("1")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestsController())(context)) {
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
		context.SetPath("/requests")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestsController())(context)) {
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
	t.Run("success get all request role 1", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/?limit=1", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestsController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all request", response.Message)
		}
	})
	t.Run("success get all request role 3", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 3)

		req := httptest.NewRequest(http.MethodPost, "/?limit=1", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestsController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all request", response.Message)
		}
	})
}

// 5. get request activity user
func TestGetRequestActivity(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(0, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestActivityController())(context)) {
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
	t.Run("failef to fetch data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestActivityController())(context)) {
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
	t.Run("success get all", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/?limit=3&offset=0", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestActivityController())(context)) {
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

// 6. get request history users
func TestGetRequestHistory(t *testing.T) {
	t.Run("unauthorized access", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(0, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestHistoryController())(context)) {
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
	t.Run("failef to fetch data", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockErrorRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestHistoryController())(context)) {
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
	t.Run("success get all", func(t *testing.T) {
		e := echo.New()
		token, _ := middlewares.CreateToken(1, "asd@mail.com", 1)

		req := httptest.NewRequest(http.MethodPost, "/?limit=3&offset=0", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/requests")

		reqController := NewRequestController(mockRequestRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if assert.NoError(t, middlewares.JWTMiddleware()(reqController.GetRequestHistoryController())(context)) {
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

type mockRequestRepository struct{}

func (m mockRequestRepository) Create(entities.Request) error {
	return nil
}
func (m mockRequestRepository) GetAdmin(return_date, request_date, status, filter_date, category string, limit, offset int) ([]entities.RequestResponse, error) {
	return nil, nil
}
func (m mockRequestRepository) GetManager(return_date, request_date, status, filter_date, category string, limit, offset int) ([]entities.RequestResponse, error) {
	return nil, nil
}
func (m mockRequestRepository) GetById(int) (entities.RequestResponse, error) {
	return entities.RequestResponse{}, nil
}
func (m mockRequestRepository) Update(entities.Request, int) error {
	return nil
}
func (m mockRequestRepository) UpdateAvailQty(int, int) error {
	return nil
}
func (m mockRequestRepository) GetAvailQty(int) (entities.Request, error) {
	return entities.Request{}, nil
}
func (m mockRequestRepository) GetEmployee(id_employee int, is_history bool, limit, offset int) ([]entities.RequestResponse, error) {
	return nil, nil
}

type mockErrorRequestRepository struct{}

func (m mockErrorRequestRepository) Create(entities.Request) error {
	return fmt.Errorf("error")
}
func (m mockErrorRequestRepository) GetAdmin(return_date, request_date, status, filter_date, category string, limit, offset int) ([]entities.RequestResponse, error) {
	return nil, fmt.Errorf("error")
}
func (m mockErrorRequestRepository) GetManager(return_date, request_date, status, filter_date, category string, limit, offset int) ([]entities.RequestResponse, error) {
	return nil, fmt.Errorf("error")
}
func (m mockErrorRequestRepository) GetById(int) (entities.RequestResponse, error) {
	return entities.RequestResponse{}, fmt.Errorf("error")
}
func (m mockErrorRequestRepository) Update(entities.Request, int) error {
	return fmt.Errorf("error")
}
func (m mockErrorRequestRepository) UpdateAvailQty(int, int) error {
	return fmt.Errorf("error")
}
func (m mockErrorRequestRepository) GetAvailQty(int) (entities.Request, error) {
	return entities.Request{}, fmt.Errorf("error")
}
func (m mockErrorRequestRepository) GetEmployee(id_employee int, is_history bool, limit, offset int) ([]entities.RequestResponse, error) {
	return nil, fmt.Errorf("error")
}
