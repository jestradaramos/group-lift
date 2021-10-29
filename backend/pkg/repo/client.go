package repo

import (
	"context"

	"github.com/jestradaramos/group-lift/backend/pkg/repo/models"
)

type Repository interface {
	AddUser(ctx context.Context, user *models.User) (*models.User, error)
}
