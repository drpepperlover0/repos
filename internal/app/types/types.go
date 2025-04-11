package types

import (
	"context"

	"github.com/drpepperlover0/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id int) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Delete(ctx context.Context, id int) error
}

type Service interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id int) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Delete(ctx context.Context, id int) error
}
