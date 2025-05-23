package usecase

import (
	"context"
	"errors"
	"user-master-service/internal/domain"
	"user-master-service/internal/repository"

	"github.com/google/uuid"
)

type UserUsecase struct {
	repo repository.UserRepositoryInterface
}

func NewUserUsecase(repo repository.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Register(ctx context.Context, name, email, password, role string) (*domain.User, error) {
	id := uuid.New().String()
	user := &domain.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password, // (тут должен быть hash)
		Role:     role,
	}
	err := u.repo.CreateUser(ctx, user)
	return user, err
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	// TODO HASHING
	if user.Password != password {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (u *UserUsecase) GetProfile(ctx context.Context, id string) (*domain.User, error) {
	return u.repo.GetUserByID(ctx, id)
}
