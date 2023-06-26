package c_auth

import (
	"km-kelas-e/helpers"
	"km-kelas-e/middlewares"

	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	v *validator.Validate
}

func NewAuthHandler(v *validator.Validate) *authHandler {
	return &authHandler{v}
}

type reqAuth struct {
	username string
	password string
}

func (ah *authHandler) AuthLogin(c echo.Context) error {
	req := reqAuth{
		username: c.QueryParam("username"),
		password: c.QueryParam("password"),
	}

	err := ah.v.Struct(&req)
	if err != nil {
		return c.JSON(400, helpers.Response{
			Message: err.Error(),
		})
	}

	token, err := checkUser(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Message: "success",
		Token:   *token,
	})
}

func checkUser(req reqAuth) (*string, error) {
	if req.username != "admin" && req.password != "admin" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, helpers.Response{Message: "username or password is wrong"})
	}

	token, err := middlewares.CreateToken(1)
	if err != nil {
		log.Println(err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, helpers.Response{Message: "something is wrong with server"})
	}

	return &token, nil
}
