package model

import (
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
)

type Response struct {
	Status int
	Body   interface{}
}

func NewResponse(status int, body interface{}) *Response {
	return &Response{
		Status: status,
		Body:   body,
	}
}

func (r *Response) WriteToCtx(ctx echo.Context) error {
	_, err := easyjson.MarshalToWriter(r, ctx.Response().Writer)

	return err
}
