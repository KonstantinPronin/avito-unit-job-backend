package middleware

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/constants"
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

func ParseErrors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := next(ctx)

		if err != nil {
			switch err := err.(type) {
			case *echo.HTTPError:
				return err
			case *model.InvalidArgumentError:
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			case *model.ForbiddenError:
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			case *model.NotFoundError:
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			case *model.ConflictError:
				return echo.NewHTTPError(http.StatusConflict, err.Error())
			case *pq.Error:
				code, msg := ParsePqErrors(err)
				return echo.NewHTTPError(code, msg)
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		return nil
	}
}

func ParsePqErrors(err *pq.Error) (int, string) {
	if err.Code.Class() != constants.IntegrityConstraintViolation {
		return http.StatusInternalServerError, err.Error()
	}

	if strings.Contains(err.Message, "transactions_from") {
		return http.StatusNotFound, "payer not found"
	}

	if strings.Contains(err.Message, "transactions_to") {
		return http.StatusNotFound, "recipient not found"
	}

	if strings.Contains(err.Message, "users_balance") {
		return http.StatusConflict, "negative balance"
	}

	return http.StatusBadRequest, err.Error()
}
