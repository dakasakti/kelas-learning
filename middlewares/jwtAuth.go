package middlewares

import (
	"km-kelas-e/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWT(secret interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.ErrUnauthorized
			}

			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				return echo.ErrUnauthorized
			}

			token, err := parseToken(splitToken[1], secret)
			if err != nil || !token.Valid {
				return echo.ErrUnauthorized
			}

			c.Set("user", token)
			return next(c)
		}
	}
}

func parseToken(tokenString string, secret interface{}) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.ErrUnauthorized
		}

		return secret, nil
	})
}

func CreateToken(userID int) (string, error) {
	payload := jwt.MapClaims{
		"userid": userID,
		"role":   "admin",
		"exp":    time.Now().Add(time.Hour * 6).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(config.JWT_SECRET))
}

func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}

	return
}
