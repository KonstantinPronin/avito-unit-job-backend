package user

import "github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"

type Repository interface {
	Add(user *model.User) (*model.User, error)
	Get(uid uint64) (*model.User, error)
	Update(uid uint64, user *model.User) (*model.User, error)
	Delete(uid uint64) error
}
