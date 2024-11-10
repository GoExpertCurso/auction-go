package user_usercase

import (
	"context"

	"github.com/GoExpertCurso/auction-go/internal/entity/user_entity"
	"github.com/GoExpertCurso/auction-go/internal/internal_error"
)

type UserUsecase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUsecaseInterface interface {
	FindUserById(ctx context.Context, userId string) (*UserOutputDTO, *internal_error.InternalError)
}

func (uc *UserUsecase) FindUserById(ctx context.Context, userId string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := uc.UserRepository.FindUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
