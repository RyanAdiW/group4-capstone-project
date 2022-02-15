package route

import (
	"sirclo/project/capstone/delivery/controllers/auth"
	"sirclo/project/capstone/delivery/controllers/user"

	middlewares "sirclo/project/capstone/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	loginController *auth.AuthController,
	userController *user.UserController) {

	// login
	e.POST("/login", loginController.LoginEmailController())

	// user
	e.POST("/users", userController.CreateUserController())
	e.GET("/users", userController.GetUsersController())
	e.GET("/users/:id", userController.GetByIdController())
	e.PUT("/users/:id", userController.UpdateUserController(), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.DeleteUserController(), middlewares.JWTMiddleware())
}
