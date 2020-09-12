package repository

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"
	errs "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type User struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (u *User) Add(usr *model.User) (*model.User, error) {
	return usr, u.db.Table("job.users").Create(usr).Error
}

func (u *User) Get(uid uint64) (*model.User, error) {
	usr := new(model.User)

	if err := u.db.Table("job.users").Where("id = ?", uid).First(usr).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errs.NewNotFoundError(err.Error())
		}

		return nil, err
	}

	return usr, nil
}

func (u *User) Update(uid uint64, usr *model.User) (*model.User, error) {
	old, err := u.Get(uid)
	if err != nil {
		return nil, err
	}

	if usr.Username != "" {
		old.Username = usr.Username
	}

	return old, u.db.Table("job.users").Save(old).Error
}

func (u *User) Contains(username string) bool {
	if err := u.db.Table("job.users").
		Where("username = ?", username).First(new(User)).Error; gorm.IsRecordNotFoundError(err) {
		return false
	}

	return true
}

func NewUser(db *gorm.DB, logger *zap.Logger) user.Repository {
	return &User{
		db:     db,
		logger: logger,
	}
}
