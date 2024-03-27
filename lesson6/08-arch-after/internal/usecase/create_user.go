package usecase

import (
	"context"

	"github.com/klakovsky/users/internal/domain"
)

func (uc *Usecase) CreateUser(ctx context.Context, name string) (domain.User, error) {
	var u domain.User
	ctx, cancel := context.WithTimeout(ctx, uc.createUserTimeout)
	defer cancel()

	u, err := uc.userRepo.CreateUser(ctx, name)
	if err != nil {
		return u, err
	}

	return u, nil
}
