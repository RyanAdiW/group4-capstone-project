package request

import (
	"log"
	"net/http"

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

// 1. create request (employee)
func (rr RequestController) CreateRequestEmployee() echo.HandlerFunc {
	return func(c echo.Context) error {
		idRole, err := middlewares.GetIdRole(c)
		if err != nil || idRole != 2 {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// bind data
		var requestReq RequestFormat
		if err := c.Bind(&requestReq); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		idUser, _ := middlewares.GetId(c)
		request := entities.Request{
			Id_user:     idUser,
			Id_asset:    requestReq.Id_asset,
			Id_status:   1,
			Return_date: requestReq.Return_date,
			Description: requestReq.Descrition,
		}

		// create user to database
		err = rr.repository.Create(request)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create request"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create request"))
	}
}
