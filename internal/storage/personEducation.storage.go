package storage

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/models"
)

// NewPersonEducationStorage  constructor para PersonEducationStorage
func NewPersonEducationStorage() PersonEducationStorage {
	return &repoPersonEducation{}
}

type repoPersonEducation struct {
}

type PersonEducationStorage interface {
	Create(ctx context.Context, personEducation models.PersonEducation) (models.PersonEducation, error)
	GetEducations(ctx context.Context, uuid string) ([]models.PersonEducation, error)
}

func (*repoPersonEducation) Create(ctx context.Context, personEducation models.PersonEducation) (models.PersonEducation, error) {
	query := `INSERT INTO personeducation VALUES(?,?,?,?,?,?,?,?,?);`

	_, err := db.QueryContext(
		ctx,
		query,
		personEducation.UUID,
		personEducation.UuidPerson,
		personEducation.Country,
		personEducation.Establishment,
		personEducation.PeriodOf,
		personEducation.PeriodTo,
		personEducation.Certificate,
		personEducation.Status,
		personEducation.Grade,
	)

	if err != nil {
		return personEducation, err
	}

	return personEducation, nil
}

func (*repoPersonEducation) GetEducations(ctx context.Context, uuid string) ([]models.PersonEducation, error) {
	education := models.PersonEducation{}
	educations := []models.PersonEducation{}
	query := `SELECT * FROM personeducation where uuidPerson = ?;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return educations, err
	}

	for rows.Next() {
		err := rows.Scan(&education.UUID,
			&education.UuidPerson,
			&education.Country,
			&education.Establishment,
			&education.PeriodOf,
			&education.PeriodTo,
			&education.Certificate,
			&education.Status,
			&education.Grade)
		if err != nil {
			return educations, err
		}

		educations = append(educations, education)
	}
	return educations, nil
}
