package route

import (
	"sirclo/project/capstone/delivery/controllers/asset"
	"sirclo/project/capstone/delivery/controllers/auth"
	"sirclo/project/capstone/delivery/controllers/request"
	"sirclo/project/capstone/delivery/controllers/user"

	middlewares "sirclo/project/capstone/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	loginController *auth.AuthController,
	userController *user.UserController,
	assetController *asset.AssetController,
	requestController *request.RequestController) {

	// login
	e.POST("/login", loginController.LoginEmailController())

	// user
	e.POST("/users", userController.CreateUserController())
	// e.GET("/users", userController.GetUsersController())
	e.GET("/users/:id", userController.GetByIdController())
	// e.PUT("/users/:id", userController.UpdateUserController(), middlewares.JWTMiddleware())
	// e.DELETE("/users/:id", userController.DeleteUserController(), middlewares.JWTMiddleware())

	// asset
	e.GET("/assets", assetController.GetAssetsController())
	e.POST("/assets/add", assetController.CreateAssetController(), middlewares.JWTMiddleware())
	e.GET("assets/summary", assetController.GetSummaryAssetsController(), middlewares.JWTMiddleware())
	e.GET("assets/detail/:id", assetController.GetAssetByIdController(), middlewares.JWTMiddleware())
	e.PUT("assets/update/:id", assetController.UpdateAssetController(), middlewares.JWTMiddleware())
	e.GET("assets/usage/:id", assetController.GetHistoryUsageController(), middlewares.JWTMiddleware())

	// request
	e.POST("/requests", requestController.CreateRequestEmployee(), middlewares.JWTMiddleware())
	e.GET("/requests", requestController.GetRequestsController(), middlewares.JWTMiddleware())
	e.GET("requests/:id", requestController.GetRequestByIdController(), middlewares.JWTMiddleware())
	e.PUT("requests/:id", requestController.UpdateRequestStatus(), middlewares.JWTMiddleware())

	// employee
	e.GET("employee/activity", requestController.GetRequestActivityController(), middlewares.JWTMiddleware())
	e.GET("employee/history", requestController.GetRequestHistoryController(), middlewares.JWTMiddleware())
	e.GET("employee/request_loan/:id", requestController.GetRequestByIdController(), middlewares.JWTMiddleware())
}
