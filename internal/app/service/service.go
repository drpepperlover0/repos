package service

import (
	"context"

	"github.com/drpepperlover0/internal/app/types"
	"github.com/drpepperlover0/internal/models"
)

type Service struct {
	repo types.UserRepository
}

func New(repo types.UserRepository) types.Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, user *models.User) error {
	return s.repo.Create(ctx, user)
}

func (s *Service) GetAll(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (*models.User, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
