package repository

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type User struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (u *User) Add(user *model.User) (*model.User, error) {
	panic("implement me")
}

func (u *User) Get(uid uint64) (*model.User, error) {
	panic("implement me")
}

func (u *User) Update(uid uint64, user *model.User) (*model.User, error) {
	panic("implement me")
}

func (u *User) Delete(uid uint64) error {
	panic("implement me")
}

func NewUser(db *gorm.DB, logger *zap.Logger) user.Repository {
	return &User{
		db:     db,
		logger: logger,
	}
}
