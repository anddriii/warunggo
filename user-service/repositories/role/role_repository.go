package repository

import (
	"context"

	"github.com/anddriii/warunggo/user-service/domain/models"
)

type IRoleRepo interface {
	GetAll(ctx context.Context, search string) ([]models.Role, error)
	GetByID(ctx context.Context, id int64) (*models.Role, error)
	Create(ctx context.Context, req *models.Role) error
	Update(ctx context.Context, req *models.Role, id int64) error
	Delete(ctx context.Context, id int64) error
}
