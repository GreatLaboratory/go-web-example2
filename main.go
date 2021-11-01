package main

import (
	"awesomeProject/handlers"
	"awesomeProject/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// create echo instance
	e := echo.New()

	// set logger config
	e.Use(middlewares.LoggerConfig())

	// track all routers
	e.Use(middlewares.RountingTracker)

	// simple hello world
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	// set user router group
	userRouter := e.Group("/users")

	// authorize user router using basic auth
	userRouter.Use(middlewares.BasicAuthorization())

	// compose user routers
	userRouter.POST("/v1", handlers.CreateUserHandler1)
	userRouter.POST("/v2", handlers.CreateUserHandler2)
	userRouter.POST("/upload/profile-img", handlers.UploadProfileImgHandler)
	userRouter.GET("/:id", handlers.GetUserHandler)
	userRouter.GET("", handlers.SearchUserHandler)
	/**
	userRouter.PUT("/:id")
	userRouter.DELETE("/:id")
	 */

	// set port
	e.Logger.Fatal(e.Start(":1323"))
}