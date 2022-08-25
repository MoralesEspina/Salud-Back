package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

// NewReferencesStorage  constructor para CurriculumStorage
func NewReferencesStorage() ReferencesStorage {
	return &repoReferences{}
}

type repoReferences struct {
}

type ReferencesStorage interface {
	Create(ctx context.Context, references models.References) (models.References, error)
	GetOne(ctx context.Context, uuid string) (models.References, error)
	GetReferences(ctx context.Context, uuid string) ([]models.References, error)
	DeleteReferences(ctx context.Context, uuid string) (string, error)
}

func (*repoReferences) Create(ctx context.Context, references models.References) (models.References, error) {
	query := `INSERT INTO u1ntiesb2kvna45k.references VALUES(?,?,?,?,?,?,?,?,?);`

	_, err := db.QueryContext(
		ctx,
		query,
		references.UUID,
		references.UuidPerson,
		references.Name,
		references.Phone,
		references.Relationship,
		references.BornDate,
		references.Profession,
		references.Company,
		references.IsFamiliar,
	)

	if err != nil {
		return references, err
	}

	return references, nil
}

func (*repoReferences) GetOne(ctx context.Context, uuid string) (models.References, error) {
	references := models.References{}

	query := `
	SELECT * FROM u1ntiesb2kvna45k.references where uuidPerson = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&references.UUID,
		&references.UuidPerson,
		&references.Name,
		&references.Phone,
		&references.Relationship,
		&references.BornDate,
		&references.Profession,
		&references.Company,
		&references.IsFamiliar,
	)

	if err == sql.ErrNoRows {
		return references, lib.ErrNotFound
	}

	if err != nil {
		return references, err
	}

	return references, nil
}

func (*repoReferences) GetReferences(ctx context.Context, uuid string) ([]models.References, error) {
	reference := models.References{}
	references := []models.References{}
	query := `SELECT * FROM u1ntiesb2kvna45k.references where uuidPerson = ?;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return references, err
	}

	for rows.Next() {
		err := rows.Scan(&reference.UUID,
			&reference.UuidPerson,
			&reference.Name,
			&reference.Phone,
			&reference.Relationship,
			&reference.BornDate,
			&reference.Profession,
			&reference.Company,
			&reference.IsFamiliar)
		if err != nil {
			return references, err
		}

		references = append(references, reference)
	}
	return references, nil
}

func (*repoReferences) DeleteReferences(ctx context.Context, uuid string) (string, error) {
	queryUpdate := "DELETE FROM u1ntiesb2kvna45k.references WHERE uuid = ?;"

	rows, err := db.ExecContext(ctx, queryUpdate, uuid)
	if err != nil {
		return "", err
	}

	resultDelete, _ := rows.RowsAffected()
	if resultDelete == 0 {
		return "", lib.ErrNotFound
	}

	return uuid, nil
}
