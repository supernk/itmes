package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// UserUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// UserRepo is a Greater repo.
type UserRepo interface {

}

// NewUserUsecase new a Greeter usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}
