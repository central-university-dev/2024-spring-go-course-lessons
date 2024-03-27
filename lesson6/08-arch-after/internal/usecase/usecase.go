package usecase

import (
	"context"
	"time"

	"github.com/klakovsky/users/internal/domain"
)

type UserRepo interface {
	GetUser(ctx context.Context, ID int) (domain.User, error)
	CreateUser(ctx context.Context, name string) (domain.User, error)
}

type Usecase struct {
	userRepo          UserRepo
	getUserTimeout    time.Duration
	createUserTimeout time.Duration
}

func New(repo UserRepo, getUserTimeout, createUserTimeout time.Duration) *Usecase {
	return &Usecase{
		userRepo:          repo,
		getUserTimeout:    getUserTimeout,
		createUserTimeout: createUserTimeout,
	}
}
