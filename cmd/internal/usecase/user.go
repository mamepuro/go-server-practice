package usecase

import (
	"context"
	"go-server-practice/cmd/internal/domain"

	"github.com/google/uuid"
)

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return u.repo.FindAll(ctx)
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return u.repo.DeleteUser(ctx, id)
}
