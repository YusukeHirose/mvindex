package handler

import (
	"fmt"
	"mvindex/service"
	"net/http"

	"github.com/labstack/echo"
)

func GetMovies(c echo.Context) error {
	fmt.Println("get movies")
	result := service.GetTopRatedList()
	return c.JSON(http.StatusOK, result)
}
