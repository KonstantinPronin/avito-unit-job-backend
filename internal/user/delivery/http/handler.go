package http

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/middleware"
	response "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type UserHandler struct {
	usecase   user.Usecase
	cur       currency.Usecase
	server    *echo.Echo
	logger    *zap.Logger
	sanitizer *bluemonday.Policy
}

func NewUserHandler(
	usecase user.Usecase,
	cur currency.Usecase,
	server *echo.Echo,
	logger *zap.Logger,
	sanitizer *bluemonday.Policy) {
	handler := &UserHandler{
		usecase:   usecase,
		cur:       cur,
		server:    server,
		logger:    logger,
		sanitizer: sanitizer,
	}

	server.POST("/user", handler.AddUser, middleware.ParseErrors)
	server.GET("/user/:uid", handler.GetUser, middleware.ParseErrors)
	server.PUT("/user/:uid", handler.UpdateUser, middleware.ParseErrors)
}

func (uh *UserHandler) AddUser(ctx echo.Context) error {
	usr, err := uh.unmarshalFromReader(ctx.Request().Body)
	if err != nil {
		return err
	}

	usr, err = uh.usecase.Add(usr)
	if err != nil {
		return err
	}

	return response.NewResponse(http.StatusOK, usr).WriteToCtx(ctx)
}

func (uh *UserHandler) GetUser(ctx echo.Context) error {
	uid, err := middleware.GetIdParam(ctx, "uid")
	if err != nil {
		return err
	}

	usr, err := uh.usecase.Get(uid)
	if err != nil {
		return err
	}

	cur := ctx.QueryParam("currency")
	if cur != "" {
		val, err := uh.cur.Exchange(usr.Balance, cur)

		if err != nil {
			return err
		}

		usr.Balance = val
	}

	return response.NewResponse(http.StatusOK, usr).WriteToCtx(ctx)
}

func (uh *UserHandler) UpdateUser(ctx echo.Context) error {
	uid, err := middleware.GetIdParam(ctx, "uid")
	if err != nil {
		return err
	}

	usr, err := uh.unmarshalFromReader(ctx.Request().Body)
	if err != nil {
		return err
	}

	usr, err = uh.usecase.Update(uid, usr)
	if err != nil {
		return err
	}

	return response.NewResponse(http.StatusOK, usr).WriteToCtx(ctx)
}

func (uh *UserHandler) sanitize(usr *model.User) {
	usr.Username = uh.sanitizer.Sanitize(usr.Username)
}

func (uh *UserHandler) unmarshalFromReader(r io.Reader) (*model.User, error) {
	usr := new(model.User)

	if err := easyjson.UnmarshalFromReader(r, usr); err != nil {
		return nil, response.NewInvalidArgument("wrong request body format")
	}
	uh.sanitize(usr)

	return usr, nil
}
