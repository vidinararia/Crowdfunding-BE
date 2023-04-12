package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func HttpParamsInt(e echo.Context, key string) int {
	intStr := e.Param(key)

	int64Var, err := strconv.ParseInt(intStr, 10, 32)
	if err != nil {
		return 0
	}
	return int(int64Var)
}
