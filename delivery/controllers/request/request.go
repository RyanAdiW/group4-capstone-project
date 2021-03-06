package request

import (
	"log"
	"net/http"
	"strconv"

	response "sirclo/project/capstone/delivery/common"
	middlewares "sirclo/project/capstone/delivery/middleware"
	"sirclo/project/capstone/entities"
	requestRepo "sirclo/project/capstone/repository/request"

	"github.com/labstack/echo/v4"
)

type RequestController struct {
	repository requestRepo.RequestRepo
}

func NewRequestController(request requestRepo.RequestRepo) *RequestController {
	return &RequestController{repository: request}
}

// 1. create request
func (rc RequestController) CreateRequestEmployee() echo.HandlerFunc {
	return func(c echo.Context) error {
		idRole, err := middlewares.GetIdRole(c)
		if err != nil || idRole == 3 {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// bind data
		var requestReq RequestFormat
		if err := c.Bind(&requestReq); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		var request entities.Request

		if idRole == 2 {
			idUser, _ := middlewares.GetId(c)
			request = entities.Request{
				Id_user:     idUser,
				Id_asset:    requestReq.Id_asset,
				Id_status:   1,
				Return_date: requestReq.Return_date,
				Description: requestReq.Descrition,
			}
		}

		if idRole == 1 {
			request = entities.Request{
				Id_user:     requestReq.Id_user,
				Id_asset:    requestReq.Id_asset,
				Id_status:   2,
				Return_date: requestReq.Return_date,
				Description: requestReq.Descrition,
			}
		}

		// create request to database
		err = rc.repository.Create(request)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create request"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create request"))
	}
}

// 2. get request details by id
func (rc RequestController) GetRequestByIdController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetIdRole(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		requestId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		request, err := rc.repository.GetById(requestId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get asset", request))
	}
}

// 3. update status request
func (rc RequestController) UpdateRequestStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		idRole, err := middlewares.GetIdRole(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		idRequest, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		// binding data
		request := entities.Request{}

		if errBind := c.Bind(&request); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		/*
			status check id
			1: menunggu persetujuan admin
			2: menunggu persetujuan manager
			3: disetujui manager
			4: ditolak manager
			5: ditolak admin
			6: diterima
			7: minta dikembalikan
			8: berhasil dikembalikan
		*/

		/*
			role
			1: admin
			2: employee
			3: manager
		*/

		switch idRole {
		// admin
		case 1:
			if request.Id_status != 2 {
				if request.Id_status != 5 {
					if request.Id_status != 6 {
						if request.Id_status != 7 {
							return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "id_status must be 2 || 5 || 6 || 7"))
						}
					}
				}
			}
		// employee
		case 2:
			if request.Id_status != 8 {
				return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "id_status must be 8"))
			}
		// manager
		case 3:
			if request.Id_status != 3 {
				if request.Id_status != 4 {
					return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "id_status must be 3 || 4"))
				}
			}
		}
		if request.Id_status == 8 || request.Id_status == 4 || request.Id_status == 5 {
			request.Return_date = "0000-00-00"
		}

		// update request based on id to database
		errUpdate := rc.repository.Update(request, idRequest)
		if errUpdate != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", errUpdate.Error()))
		}

		// if status == 6 (diterima)
		if request.Id_status == 6 {
			// get current avail qty
			availQty, _ := rc.repository.GetAvailQty(idRequest)
			if availQty.Avail_quantity >= 0 {
				qty := availQty.Avail_quantity - 1
				err := rc.repository.UpdateAvailQty(qty, idRequest)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", errUpdate.Error()))
				}
			}
		}
		// if status == 8 berhasil dikembalikan
		if request.Id_status == 8 {
			// get current avail qty
			availQty, _ := rc.repository.GetAvailQty(idRequest)
			if availQty.Avail_quantity < availQty.Initial_quantity {
				qty := availQty.Avail_quantity + 1
				err := rc.repository.UpdateAvailQty(qty, idRequest)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", errUpdate.Error()))
				}
			}
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update request"))
	}
}

// 4. get requests
func (rc *RequestController) GetRequestsController() echo.HandlerFunc {
	return func(c echo.Context) error {
		idRole, err := middlewares.GetIdRole(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		returnDate := c.QueryParam("return_date")
		requestDate := c.QueryParam("request_date")
		status := c.QueryParam("status")
		filterDate := c.QueryParam("filter_date")
		category := c.QueryParam("category")
		limitStr := c.QueryParam("limit")
		offsetStr := c.QueryParam("offset")

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 0
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}

		var requests []entities.RequestResponse
		totalPage := 0
		switch idRole {
		case 1:
			requests, err = rc.repository.GetAdmin(returnDate, requestDate, status, filterDate, category, limit, offset)
			if limit > 0 {
				requestsTotPage, _ := rc.repository.GetAdmin(returnDate, requestDate, status, filterDate, category, 0, 0)
				totalPage = (len(requestsTotPage) / limit) + 1
			}
		case 3:
			requests, err = rc.repository.GetManager(returnDate, requestDate, status, filterDate, category, limit, offset)
			if limit > 0 {
				requestsTotPage, _ := rc.repository.GetManager(returnDate, requestDate, status, filterDate, category, 0, 0)

				if len(requestsTotPage)%limit == 0 {
					totalPage = (len(requestsTotPage) / limit)
				} else {
					totalPage = (len(requestsTotPage) / limit) + 1
				}
			}
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all request", map[string]interface{}{
			"total_page": totalPage,
			"data":       requests,
		}))
	}
}

// get request activity users
func (rc RequestController) GetRequestActivityController() echo.HandlerFunc {
	return func(c echo.Context) error {
		idUser, err := middlewares.GetId(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		limitStr := c.QueryParam("limit")
		offsetStr := c.QueryParam("offset")

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 0
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}

		request, err := rc.repository.GetEmployee(idUser, false, limit, offset)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		totalPage := 0
		if limit > 0 {
			requestsTotPage, _ := rc.repository.GetEmployee(idUser, false, 0, 0)

			if len(requestsTotPage)%limit == 0 {
				totalPage = (len(requestsTotPage) / limit)
			} else {
				totalPage = (len(requestsTotPage) / limit) + 1
			}
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all assets", map[string]interface{}{
			"total_page": totalPage,
			"data":       request,
		}))
	}
}

// get request history users
func (rc RequestController) GetRequestHistoryController() echo.HandlerFunc {
	return func(c echo.Context) error {
		idUser, err := middlewares.GetId(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		limitStr := c.QueryParam("limit")
		offsetStr := c.QueryParam("offset")

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 0
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}

		request, err := rc.repository.GetEmployee(idUser, true, limit, offset)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		totalPage := 0
		if limit > 0 {
			requestsTotPage, _ := rc.repository.GetEmployee(idUser, true, 0, 0)

			if len(requestsTotPage)%limit == 0 {
				totalPage = (len(requestsTotPage) / limit)
			} else {
				totalPage = (len(requestsTotPage) / limit) + 1
			}
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all assets", map[string]interface{}{
			"total_page": totalPage,
			"data":       request,
		}))
	}
}
