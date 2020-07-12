package handler

import (
	"mvindex/service"
	"net/http"

	"github.com/labstack/echo"
)

func GetMovies(c echo.Context) error {
	page := c.QueryParam("page")
	result := service.GetTopRatedList(page)
	return c.JSON(http.StatusOK, result)
}

func GetMovieDetail(c echo.Context) error {
	id := c.Param("id")
	result := service.GetDetail(id)
	return c.
		JSON(http.StatusOK, result)
}

func GetMovieByKeyword(c echo.Context) error {
	kw := c.QueryParam("keyword")
	return c.JSON(http.StatusOK, kw)
}
