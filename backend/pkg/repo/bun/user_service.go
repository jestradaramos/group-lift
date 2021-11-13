package bun

import (
	"context"

	"github.com/jestradaramos/group-lift/backend/pkg/repo/models"
)

// Need to find the right place for this
func (d *DB) CreateUserTable(ctx context.Context) {
	_, err := d.DB.NewCreateTable().
		Model((*models.User)(nil)).
		Exec(ctx)
	if err != nil {
		panic(err)
	}

}

func (r *DB) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := r.DB.NewInsert().
		Model(user).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil

}
