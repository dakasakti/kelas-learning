package middlewares

import (
	"km-kelas-e/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BasicAuth(username, password string, secret interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.QueryParam("username")
			pass := c.QueryParam("password")

			if user != username || pass != password {
				return echo.ErrUnauthorized
			}

			tokenString, err := CreateToken(1)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, helpers.Response{Message: "something is wrong with server"})
			}

			token, err := parseToken(tokenString, secret)
			if err != nil || !token.Valid {
				return echo.ErrUnauthorized
			}

			c.Set("user", token)
			return next(c)
		}
	}
}
