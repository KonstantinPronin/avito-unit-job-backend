package user

import "github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"

type Repository interface {
	Add(usr *model.User) (*model.User, error)
	Get(uid uint64) (*model.User, error)
	Update(uid uint64, usr *model.User) (*model.User, error)
	Contains(username string) bool
}
