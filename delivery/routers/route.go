package route

import (
	"sirclo/project/capstone/delivery/controllers/user"

	// middlewares "sirclo/groupproject/restapi/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	userController *user.UserController) {

	// login
	// e.POST("/login", loginController.LoginUserNameController())

	// user
	e.POST("/users", userController.CreateUserController())
	e.GET("/users", userController.GetUsersController())
	e.GET("/users/:id", userController.GetByIdController())
	e.PUT("/users/:id", userController.UpdateUserController())
	e.DELETE("/users/:id", userController.DeleteUserController())
}
