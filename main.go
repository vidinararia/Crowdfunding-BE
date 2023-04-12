package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vidinararia/crowd-funding-be/router"
)

func main() {
	ec := echo.New()

	router.InitRouter(ec)
	ec.Start("localhost:8080")
}
