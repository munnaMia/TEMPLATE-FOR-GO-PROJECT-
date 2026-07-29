package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/munnaMia/ahlan/internal/domain"
	"github.com/munnaMia/ahlan/internal/infra/auth"
)

type UserUsecase struct {
	repo     domain.UserRepository
	tokenSvr domain.TokenService
}

func NewUserUsecase(r domain.UserRepository, tknSvr domain.TokenService) *UserUsecase {
	return &UserUsecase{
		repo:     r,
		tokenSvr: tknSvr,
	}
}

type RegisterResult struct {
	user  *domain.User
	token string
}

func (uc *UserUsecase) Register(ctx context.Context, name, email, password string) (*RegisterResult, error) {

	// check user or the email already exist or not first...

	user, err := uc.repo.Create(ctx, &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	token, err := uc.tokenSvr.GenerateToken(auth.Payload{
		Sub:      strconv.Itoa(user.Id),
		UserName: user.Name,
		Email:    user.Email,
		IAT:      time.Now().Unix(),
		EXP:      time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		user:  user,
		token: token,
	}, nil
}
