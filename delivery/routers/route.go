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
	e.GET("/users", userController.GetUsersController())
	e.GET("/users/:id", userController.GetByIdController())
	e.PUT("/users/:id", userController.UpdateUserController(), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.DeleteUserController(), middlewares.JWTMiddleware())

	// asset
	e.POST("/assets", assetController.CreateAssetController(), middlewares.JWTMiddleware())
	e.GET("/assets", assetController.GetAssetsController())
	e.GET("assets/summary", assetController.GetSummaryAssetsController(), middlewares.JWTMiddleware())
	e.GET("assets/:id", assetController.GetAssetByIdController(), middlewares.JWTMiddleware())
	e.GET("assets/history/:id", assetController.GetHistoryUsageController(), middlewares.JWTMiddleware())
	e.PUT("assets/:id", assetController.UpdateAssetController(), middlewares.JWTMiddleware())
	e.DELETE("assets/:id", assetController.DeleteAssetController(), middlewares.JWTMiddleware())

	// request
	e.POST("/requests", requestController.CreateRequestEmployee(), middlewares.JWTMiddleware())
	e.GET("requests/:id", requestController.GetRequestByIdController(), middlewares.JWTMiddleware())
}
