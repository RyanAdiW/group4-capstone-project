package asset

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"sirclo/project/capstone/entities"
	"sirclo/project/capstone/util"

	response "sirclo/project/capstone/delivery/common"
	middlewares "sirclo/project/capstone/delivery/middleware"
	assetRepo "sirclo/project/capstone/repository/asset"

	"github.com/labstack/echo/v4"
)

type AssetController struct {
	repository assetRepo.AssetRepo
}

func NewAssetController(asset assetRepo.AssetRepo) *AssetController {
	return &AssetController{repository: asset}
}

// 1. create asset controller
func (ac AssetController) CreateAssetController() echo.HandlerFunc {
	return func(c echo.Context) error {
		idrole, err := middlewares.GetIdRole(c)

		if err != nil || idrole != 1 {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// bind data
		var userRequest UserRequestFormat
		if err := c.Bind(&userRequest); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		//bind data photo
		// Multipart form
		var url_photo string
		form, err := c.MultipartForm()
		if err == nil {
			files := form.File["photo"]

			for _, file := range files {
				// Source
				src, err := file.Open()
				if err != nil {
					log.Println(err)
					return err
				}
				defer src.Close()

				fileExtension := filepath.Ext(file.Filename)
				err = util.CheckExtension(fileExtension)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to join, photo format not allowed"))
				}

				fileSize := file.Size
				err = util.CheckSize(fileSize)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to join, photo size too big"))
				}

				filename := "assets_pic/" + strconv.Itoa(userRequest.Id_category) + "_" + userRequest.Name + fileExtension
				url_photo, err = util.UploadToS3(&src, filename)
				if err != nil {
					return err
				}
			}
		}

		asset := entities.Asset{
			Name:             userRequest.Name,
			Description:      userRequest.Description,
			Initial_quantity: userRequest.Initial_quantity,
			Avail_quantity:   userRequest.Initial_quantity,
			Photo:            url_photo,
			Id_category:      userRequest.Id_category,
		}

		// create user to database
		err = ac.repository.Create(asset)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create asset"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create asset"))
	}
}

// 2. get all user controller
func (ac AssetController) GetAssetsController() echo.HandlerFunc {
	return func(c echo.Context) error {
		category := c.QueryParam("category")
		maintenance := c.QueryParam("maintenance")
		avail := c.QueryParam("avail")

		limitstr := c.QueryParam("limit")
		offsetstr := c.QueryParam("offset")

		limit, err := strconv.Atoi(limitstr)
		if err != nil {
			limit = 0
		}

		offset, err := strconv.Atoi(offsetstr)
		if err != nil {
			offset = 0
		}

		assets, err := ac.repository.Get(category, maintenance, avail, limit, offset)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		total_page := 0
		if limit > 0 {
			assetsfortotpage, _ := ac.repository.Get(category, maintenance, avail, 0, 0)
			if len(assetsfortotpage)%limit == 0 {
				total_page = (len(assetsfortotpage) / limit)
			} else {
				total_page = (len(assetsfortotpage) / limit) + 1
			}
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all assets", map[string]interface{}{
			"total_page": total_page,
			"data":       assets,
		}))
	}
}

// 3. get asset by id
func (ac AssetController) GetAssetByIdController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetIdRole(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		assetid, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// get user from db
		asset, err := ac.repository.GetById(assetid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get asset", asset))
	}
}

// 4. update asset
func (ac AssetController) UpdateAssetController() echo.HandlerFunc {
	return func(c echo.Context) error {
		idrole, err := middlewares.GetIdRole(c)

		if err != nil || idrole != 1 {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		id_asset, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		// binding data
		asset := entities.Asset{}
		if errBind := c.Bind(&asset); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		//asset existed
		assetExisted, err := ac.repository.GetById(id_asset)
		if err != nil {
			log.Println("err get date asset: ", err)
			return c.JSON(http.StatusInternalServerError, response.InternalServerError("error", "err get date asset"))
		}

		//bind data photo
		// Multipart form
		var url_photo string
		form, err := c.MultipartForm()
		name := asset.Name
		if name == "" {
			name = assetExisted.Name
		}
		if err == nil {
			files := form.File["photo"]

			for _, file := range files {
				// Source
				src, err := file.Open()
				if err != nil {
					log.Println(err)
					return err
				}
				defer src.Close()

				fileExtension := filepath.Ext(file.Filename)
				err = util.CheckExtension(fileExtension)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to join, photo format not allowed"))
				}

				fileSize := file.Size
				err = util.CheckSize(fileSize)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to join, photo size too big"))
				}

				filename := "assets_pic/" + strconv.Itoa(asset.Id_category) + "_" + name + fileExtension
				url_photo, err = util.UploadToS3(&src, filename)
				if err != nil {
					return err
				}
			}
		}
		if url_photo != "" {
			asset.Photo = url_photo
		}

		// update user based on id to database
		errUpdate := ac.repository.Update(assetExisted, asset, id_asset)
		if errUpdate != nil {
			fmt.Println(errUpdate)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", errUpdate.Error()))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update asset"))
	}
}

// 5. delete asset
func (ac AssetController) DeleteAssetController() echo.HandlerFunc {
	return func(c echo.Context) error {
		idrole, err := middlewares.GetIdRole(c)

		if err != nil || idrole != 1 {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		id_asset, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		// delete asset based on id from database
		errDelete := ac.repository.Delete(id_asset)
		if errDelete != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "delete success"))
	}
}

// 6. get summary asset
func (ac AssetController) GetSummaryAssetsController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetIdRole(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		summary, err := ac.repository.GetSummaryAsset()
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get summary assets", summary))
	}
}

// 7. get history usage
func (ac AssetController) GetHistoryUsageController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetIdRole(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		id_asset, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			fmt.Println(errConv)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}

		limitstr := c.QueryParam("limit")
		offsetstr := c.QueryParam("offset")

		limit, err := strconv.Atoi(limitstr)
		if err != nil {
			limit = 0
		}

		offset, err := strconv.Atoi(offsetstr)
		if err != nil {
			offset = 0
		}

		history, err := ac.repository.GetHistoryUsage(id_asset, limit, offset)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get history", history))
	}
}

func (ac AssetController) GetCategoriesController() echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := ac.repository.GetCategory()
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get categories", categories))
	}
}
