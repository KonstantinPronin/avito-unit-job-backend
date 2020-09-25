package usecase

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/mock"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var (
	test = model.User{
		ID:       1,
		Username: "test",
		Balance:  10,
	}

	wrongID = uint64(0)
)

func setUp(t *testing.T) (*mock.MockRepository, user.Usecase) {
	ctrl := gomock.NewController(t)
	r := mock.NewMockRepository(ctrl)
	u := NewUser(r, zap.NewExample())

	return r, u
}

func TestUser_Add_EmptyUsername(t *testing.T) {
	_, u := setUp(t)

	_, err := u.Add(new(model.User))

	assert.NotNil(t, err)
}

func TestUser_Add_ExistingUsername(t *testing.T) {
	r, u := setUp(t)

	r.EXPECT().Contains(test.Username).Return(true)

	_, err := u.Add(&model.User{Username: test.Username})

	assert.NotNil(t, err)
}

func TestUser_Add_CorrectUsername(t *testing.T) {
	r, u := setUp(t)

	r.EXPECT().Contains(test.Username).Return(false)
	r.EXPECT().Add(&test).Return(&test, nil)

	usr, err := u.Add(&test)

	assert.Nil(t, err)
	assert.Equal(t, &test, usr)
}

func TestUser_Get_WrongId(t *testing.T) {
	_, u := setUp(t)

	_, err := u.Get(wrongID)

	assert.NotNil(t, err)
}

func TestUser_Get_CorrectId(t *testing.T) {
	r, u := setUp(t)

	r.EXPECT().Get(test.ID).Return(&test, nil)

	usr, err := u.Get(test.ID)

	assert.Nil(t, err)
	assert.Equal(t, &test, usr)
}

func TestUser_Update_WrongId(t *testing.T) {
	_, u := setUp(t)

	_, err := u.Update(wrongID, nil)

	assert.NotNil(t, err)
}

func TestUser_Update_CorrectId(t *testing.T) {
	r, u := setUp(t)

	r.EXPECT().Update(test.ID, &test).Return(&test, nil)

	usr, err := u.Update(test.ID, &test)

	assert.Nil(t, err)
	assert.Equal(t, &test, usr)
}
