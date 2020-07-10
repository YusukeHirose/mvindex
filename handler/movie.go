package handler

import (
	"fmt"
	"mvindex/service"

	"github.com/labstack/echo"
)

func GetMovies(c echo.Context) error {
	fmt.Println("get movies")
	service.GetTopRatedList()
	return nil
}
