package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	log.Print("just debug")
	return c.String(http.StatusOK, "id : "+id)
}

func SearchUserHandler(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	return c.String(http.StatusOK, "name : "+name+" // age : "+age)
}

func CreateUserHandler1(c echo.Context) error {
	name := c.FormValue("name")
	age := c.FormValue("age")
	email := c.FormValue("email")
	return c.String(http.StatusCreated, "name : "+name+" // age : "+age+" // email : "+email)
}

// bind json request body into Go struct based on "Content-Type" request header
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func CreateUserHandler2(c echo.Context) error {
	user := new(User)
	err := c.Bind(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func UploadProfileImgHandler(c echo.Context) error {
	// get file from http request
	profile, err := c.FormFile("profile")
	if err != nil {
		return err
	}

	// get source
	source, err := profile.Open()
	if err != nil {
		return err
	}
	defer func(source multipart.File) {
		err := source.Close()
		if err != nil {

		}
	}(source)

	// set file destination -> project root directory
	dst, err := os.Create(profile.Filename)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	// copy file in dst
	_, err = io.Copy(dst, source)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "upload success")
}
