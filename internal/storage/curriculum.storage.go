package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

// NewCurriculumStorage  constructor para CurriculumStorage
func NewCurriculumStorage() CurriculumStorage {
	return &repoCurriculum{}
}

type repoCurriculum struct {
}

type CurriculumStorage interface {
	Create(ctx context.Context, curriculum models.Curriculum) (models.Curriculum, error)
	GetOne(ctx context.Context, uuid string) (models.Curriculum, error)
}

func (*repoCurriculum) Create(ctx context.Context, curriculum models.Curriculum) (models.Curriculum, error) {
	query := `INSERT INTO curriculum VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	_, err := db.QueryContext(
		ctx,
		query,
		curriculum.UUID,
		curriculum.UuidPerson,
		curriculum.Direction,
		curriculum.Country,
		curriculum.HomePhone,
		curriculum.BornPlace,
		curriculum.Nacionality,
		curriculum.Municipality,
		curriculum.Village,
		curriculum.WorkPhone,
		curriculum.Age,
		curriculum.CivilStatus,
		curriculum.Etnia,
		curriculum.Passport,
		curriculum.License,
	)

	if err != nil {
		return curriculum, err
	}

	return curriculum, nil
}

func (*repoCurriculum) GetOne(ctx context.Context, uuid string) (models.Curriculum, error) {
	curriculum := models.Curriculum{}

	query := `
	SELECT * FROM curriculum;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&curriculum.UUID,
		&curriculum.UuidPerson,
		&curriculum.Direction,
		&curriculum.Country,
		&curriculum.HomePhone,
		&curriculum.BornPlace,
		&curriculum.Nacionality,
		&curriculum.Municipality,
		&curriculum.Village,
		&curriculum.WorkPhone,
		&curriculum.Age,
		&curriculum.CivilStatus,
		&curriculum.Etnia,
		&curriculum.Passport,
		&curriculum.License,
	)

	if err == sql.ErrNoRows {
		return curriculum, lib.ErrNotFound
	}

	if err != nil {
		return curriculum, err
	}

	return curriculum, nil
}
