package http

import (
	"bytes"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/mock"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"
	mock2 "github.com/KonstantinPronin/avito-unit-job-backend/pkg/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"testing"
)

var (
	test = model.User{
		ID:       1,
		Username: "test",
		Balance:  10,
	}

	correctID = "1"
	wrongID   = "test"
)

func beforeTest(t *testing.T) (
	*mock.MockUsecase,
	*mock2.MockContext,
	*UserHandler) {
	ctrl := gomock.NewController(t)

	usecase := mock.NewMockUsecase(ctrl)
	handler := &UserHandler{usecase: usecase, logger: zap.NewExample(), sanitizer: bluemonday.UGCPolicy()}

	ctx := mock2.NewMockContext(ctrl)
	writer := mock2.NewMockResponseWriter(ctrl)

	response := echo.NewResponse(writer, echo.New())
	ctx.EXPECT().Response().Return(response).AnyTimes()
	writer.EXPECT().Write(gomock.Any())

	return usecase, ctx, handler
}

func TestUserHandler_AddUser_EmptyBody(t *testing.T) {
	_, ctx, handler := beforeTest(t)

	request, err := http.NewRequest("", "", strings.NewReader(""))
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	ctx.EXPECT().Request().Return(request).AnyTimes()

	err = handler.AddUser(ctx)

	assert.NotNil(t, err)
}

func TestUserHandler_AddUser(t *testing.T) {
	usecase, ctx, handler := beforeTest(t)

	body, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	request, err := http.NewRequest("", "", bytes.NewReader(body))
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	ctx.EXPECT().Request().Return(request).AnyTimes()
	usecase.EXPECT().Add(&test).Return(&test, nil)

	err = handler.AddUser(ctx)

	assert.Nil(t, err)
}

func TestUserHandler_GetUser_WrongId(t *testing.T) {
	_, ctx, handler := beforeTest(t)

	ctx.EXPECT().Param(gomock.Any()).Return(wrongID)

	err := handler.GetUser(ctx)

	assert.NotNil(t, err)
}

func TestUserHandler_GetUser(t *testing.T) {
	usecase, ctx, handler := beforeTest(t)

	ctx.EXPECT().Param(gomock.Any()).Return(correctID)
	ctx.EXPECT().QueryParam("currency").Return("")
	usecase.EXPECT().Get(test.ID).Return(&test, nil)

	err := handler.GetUser(ctx)

	assert.Nil(t, err)
}

func TestUserHandler_UpdateUser(t *testing.T) {
	usecase, ctx, handler := beforeTest(t)

	body, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	request, err := http.NewRequest("", "", bytes.NewReader(body))
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	ctx.EXPECT().Param(gomock.Any()).Return(correctID)
	ctx.EXPECT().Request().Return(request).AnyTimes()
	usecase.EXPECT().Update(test.ID, &test).Return(&test, nil)

	err = handler.UpdateUser(ctx)

	assert.Nil(t, err)
}
