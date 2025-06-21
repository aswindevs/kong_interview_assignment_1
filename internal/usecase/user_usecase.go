package usecase

import (
	"context"
	"fmt"

	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
	appError "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase/repo"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(r repo.UserRepo) *UserUsecase {
	return &UserUsecase{
		UserRepo: r,
	}
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	user, err := uc.UserRepo.FindById(ctx, id)
	if err != nil {
		return nil, appError.NewNotFoundError(fmt.Sprintf("User with id %v , not found", id))
	}
	return user, nil
}
