package usecase

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"
	errs "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"go.uber.org/zap"
)

type User struct {
	rep    user.Repository
	logger *zap.Logger
}

func (u *User) Add(usr *model.User) (*model.User, error) {
	if usr.Username == "" {
		return nil, errs.NewInvalidArgument("empty username")
	}

	if u.rep.Contains(usr.Username) {
		return nil, errs.NewConflictError("username already exist")
	}

	return u.rep.Add(usr)
}

func (u *User) Get(uid uint64) (*model.User, error) {
	if uid == 0 {
		return nil, errs.NewInvalidArgument("wrong user id")
	}

	return u.rep.Get(uid)
}

func (u *User) Update(uid uint64, usr *model.User) (*model.User, error) {
	if uid == 0 {
		return nil, errs.NewInvalidArgument("wrong user id")
	}

	return u.rep.Update(uid, usr)
}

func NewUser(rep user.Repository, logger *zap.Logger) user.Usecase {
	return &User{
		rep:    rep,
		logger: logger,
	}
}
