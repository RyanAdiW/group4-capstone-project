package main

import (
	"log"
	"sirclo/project/capstone/config"
	_route "sirclo/project/capstone/delivery/routers"
	"sirclo/project/capstone/util"

	_assetController "sirclo/project/capstone/delivery/controllers/asset"
	_authController "sirclo/project/capstone/delivery/controllers/auth"
	_userController "sirclo/project/capstone/delivery/controllers/user"

	_assetRepo "sirclo/project/capstone/repository/asset"
	_authRepo "sirclo/project/capstone/repository/auth"
	_userRepo "sirclo/project/capstone/repository/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//log
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// load config
	config := config.GetConfig()
	// initialize database connection
	db := util.MysqlDriver(config)
	defer db.Close()

	// initialize model
	authRepo := _authRepo.NewAuthRepo(db)
	userRepo := _userRepo.NewUserRepo(db)
	assetRepo := _assetRepo.NewAssetRepo(db)

	// initialize controller
	authController := _authController.NewAuthController(authRepo)
	userController := _userController.NewUserController(userRepo)
	assetController := _assetController.NewAssetController(assetRepo)

	// create new echo
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS())

	_route.RegisterPath(e, authController, userController, assetController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":80"))
}
