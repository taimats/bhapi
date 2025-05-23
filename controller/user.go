package controller

import (
	"context"

	"github.com/taimats/bhapi/domain"
	"github.com/taimats/bhapi/infra/repository"
	"github.com/taimats/bhapi/utils"
)

type User struct {
	ur *repository.User
}

func NewUser(ur *repository.User) *User {
	return &User{ur: ur}
}

func (uc *User) RegisterUser(ctx context.Context, user *domain.User) error {
	_, err := uc.ur.FindUserByAuthUserId(ctx, user.AuthUserId)
	if err == nil {
		return utils.NewErrChains(utils.ErrAlrExists, nil)
	}

	_, err = uc.ur.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *User) GetUser(ctx context.Context, authUserId string) (*domain.User, error) {
	user, err := uc.ur.FindUserByAuthUserId(ctx, authUserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *User) DeleteUser(ctx context.Context, authUserId string) error {
	user, err := uc.ur.FindUserByAuthUserId(ctx, authUserId)
	if err != nil {
		return err
	}

	err = uc.ur.DleteUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *User) UpdateUser(ctx context.Context, user *domain.User) error {
	if err := uc.ur.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
