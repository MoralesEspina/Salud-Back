package storage

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/models"
)

// NewWorkExpStorage  constructor para WorkExpStorage
func NewWorkExpStorage() WorkExpStorage {
	return &repoWorkExp{}
}

type repoWorkExp struct {
}

type WorkExpStorage interface {
	Create(ctx context.Context, workExp models.WorkExp) (models.WorkExp, error)
	GetWorks(ctx context.Context, uuid string) ([]models.WorkExp, error)
}

func (*repoWorkExp) Create(ctx context.Context, workExp models.WorkExp) (models.WorkExp, error) {
	query := `INSERT INTO workExp VALUES(?,?,?,?,?,?,?,?,?,?,?,?);`

	_, err := db.QueryContext(
		ctx,
		query,
		workExp.UUID,
		workExp.UuidPerson,
		workExp.Direction,
		workExp.Phone,
		workExp.Reason,
		workExp.DateOf,
		workExp.DateTo,
		workExp.Job,
		workExp.BossName,
		workExp.Sector,
		workExp.Salary,
		workExp.WorkExpCol,
	)

	if err != nil {
		return workExp, err
	}

	return workExp, nil
}

func (*repoWorkExp) GetWorks(ctx context.Context, uuid string) ([]models.WorkExp, error) {
	work := models.WorkExp{}
	works := []models.WorkExp{}
	query := `SELECT * FROM workexp where uuidPerson = ?;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return works, err
	}

	for rows.Next() {
		err := rows.Scan(&work.UUID,
			&work.UuidPerson,
			&work.Direction,
			&work.Phone,
			&work.Reason,
			&work.DateOf,
			&work.DateTo,
			&work.Job,
			&work.BossName,
			&work.Sector,
			&work.Salary,
			&work.WorkExpCol)
		if err != nil {
			return works, err
		}

		works = append(works, work)
	}
	return works, nil
}
