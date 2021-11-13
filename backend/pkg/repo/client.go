package repo

import (
	"context"

	"github.com/jestradaramos/group-lift/backend/pkg/repo/models"
)

type Repository interface {
	AddUser(ctx context.Context, user *models.User) (*models.User, error)
	AddLiftSession(ctx context.Context, user *models.LiftSession) (*models.LiftSession, error)
	AddLift(ctx context.Context, lift *models.Lift) (*models.Lift, error)
}
