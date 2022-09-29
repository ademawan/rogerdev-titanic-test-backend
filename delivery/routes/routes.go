package routes

import (
	"rogerdev-titanic-test-backend/delivery/controllers/auth"

	"rogerdev-titanic-test-backend/delivery/controllers/person"
	"rogerdev-titanic-test-backend/delivery/controllers/user"

	"rogerdev-titanic-test-backend/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	aa *auth.AuthController,
	uc *user.UserController,
	pc *person.PersonController,

) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())
	e.GET("google/login", aa.LoginGoogle())
	e.GET("google/callback", aa.LoginGoogleCallback())
	e.POST("users/logout", aa.Logout(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())

	//ROUTE PERSONS
	e.POST("person/post", pc.Post())
	e.GET("person/post/histories", pc.GetAll())

}
