package handler

import (
	"fmt"

	"github.com/labstack/echo"
)

func GetMovies(c echo.Context) error {
	fmt.Println("get movies")
	return nil
}
