package bun

import (
	"context"

	"github.com/jestradaramos/group-lift/backend/pkg/repo/models"
)

func (r *DB) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := r.db.NewInsert().
		Model(user).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil

}
