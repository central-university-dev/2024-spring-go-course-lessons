package usecase

import (
	"context"

	"github.com/klakovsky/users/internal/domain"
)

func (uc *Usecase) GetUser(ctx context.Context, ID int) (domain.User, error) {
	var result domain.User
	ctx, cancel := context.WithTimeout(ctx, uc.getUserTimeout)
	defer cancel()

	result, err := uc.userRepo.GetUser(ctx, ID)
	if err != nil {
		return result, err
	}

	return result, nil
}
