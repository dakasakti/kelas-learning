package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"km-kelas-e/config"
	c_articles "km-kelas-e/controller/articles"
	c_auth "km-kelas-e/controller/auth"

	"km-kelas-e/middlewares"
)

func New() *echo.Echo {

	//initiate
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddleware(e)

	ah := c_auth.NewAuthHandler(validator.New())

	auth := e.Group("auth")
	auth.POST("/login", ah.AuthLogin)

	jwtAuthGroup := e.Group("jwt")
	jwtAuthGroup.Use(middlewares.JWT([]byte(config.JWT_SECRET)))

	jwtAuthGroup.GET("/articles", c_articles.GetAllArticle)

	basicAuthGroup := e.Group("basic")
	basicAuthGroup.Use(middlewares.BasicAuth("admin", "admin"))

	basicAuthGroup.GET("/articles", c_articles.GetAllArticle)

	// router
	// article:= e.Group("article")
	// article.Use()
	// e.GET("/article/:id", getArticle)
	// e.POST("/article", addArticle)
	// e.PUT("/article/:id", updateArticle)

	return e
}
