package bun

import (
	"context"

	"github.com/jestradaramos/group-lift/backend/pkg/repo/models"
)

// Need to find the right place for this
func (d *DB) CreateLiftTables(ctx context.Context) {
	_, err := d.DB.NewCreateTable().
		Model((*models.LiftSession)(nil)).
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	_, err = d.DB.NewCreateTable().
		Model((*models.Lift)(nil)).
		ForeignKey(`("session_id") REFERENCES "lift_sessions" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if err != nil {
		panic(err)
	}

}

func (d *DB) AddLiftSession(ctx context.Context, session *models.LiftSession) (*models.LiftSession, error) {
	_, err := d.DB.NewInsert().
		Model(session).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (d *DB) AddLift(ctx context.Context, lift *models.Lift) (*models.Lift, error) {
	_, err := d.DB.NewInsert().
		Model(lift).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return lift, nil
}

func (d *DB) GetLiftSession(ctx context.Context, id string) (*models.LiftSession, error) {
	liftSession := &models.LiftSession{}
	err := d.DB.NewSelect().
		Model(liftSession).
		Relation("Lift").
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return liftSession, nil
}
