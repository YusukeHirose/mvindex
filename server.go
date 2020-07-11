package main

import (
	"mvindex/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	g := e.Group("/movies")
	g.GET("", handler.GetMovies)

	// サーバー起動
	e.Start("0.0.0.0:8080")
}