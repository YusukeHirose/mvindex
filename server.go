package main

import (
	"mvindex/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	g := e.Group("/movies")
	g.GET("", handler.GetMovies)
	g.GET("/:id", handler.GetMovieDetail)
	g.GET("/search", handler.GetMovieByKeyword)

	// サーバー起動
	e.Start("0.0.0.0:8080")
}
