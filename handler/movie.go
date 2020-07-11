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
