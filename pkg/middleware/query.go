package middleware

import (
	response "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/labstack/echo"
	"strconv"
)

func GetIdParam(ctx echo.Context, param string) (uint64, error) {
	id, err := strconv.Atoi(ctx.Param(param))

	if err != nil || id < 0 {
		return 0, response.NewInvalidArgument("wrong id format")
	}

	return uint64(id), nil
}
