package main

import (
	"rogerdev-titanic-test-backend/configs"
	ac "rogerdev-titanic-test-backend/delivery/controllers/auth"
	pc "rogerdev-titanic-test-backend/delivery/controllers/person"
	uc "rogerdev-titanic-test-backend/delivery/controllers/user"

	"rogerdev-titanic-test-backend/delivery/routes"
	authRepo "rogerdev-titanic-test-backend/repository/auth"
	personRepo "rogerdev-titanic-test-backend/repository/person"
	userRepo "rogerdev-titanic-test-backend/repository/user"

	"fmt"
	"rogerdev-titanic-test-backend/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config := configs.GetConfig()

	db, err := utils.InitDB(config)
	if err != nil {
		fmt.Println(err.Error(), "database belum terhubung, data yang di inputkan tidak akan disimpan")
	}
	defer db.Close()

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)
	personRepo := personRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)
	personController := pc.New(personRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, authController, userController, personController)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
