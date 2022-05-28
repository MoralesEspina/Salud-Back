package storage

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/models"
)

// NewReferencesStorage  constructor para CurriculumStorage
func NewReferencesStorage() ReferencesStorage {
	return &repoReferences{}
}

type repoReferences struct {
}

type ReferencesStorage interface {
	CreateRefFamiliar(ctx context.Context, references models.References) (models.References, error)
	GetRefFam(ctx context.Context, uuid string) ([]models.References, error)
	GetRefPer(ctx context.Context, uuid string) ([]models.References, error)
}

func (*repoReferences) CreateRefFamiliar(ctx context.Context, references models.References) (models.References, error) {
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

func (*repoReferences) GetRefFam(ctx context.Context, uuid string) ([]models.References, error) {
	reference := models.References{}
	references := []models.References{}

	query := `
	SELECT  uuid,
			name,
			phone,
			relationship,
			borndate FROM u1ntiesb2kvna45k.references where uuidPerson = ? And isFamiliar = true;`

	rows, err := db.QueryContext(ctx, query, uuid)

	if err != nil {
		return references, err
	}
	for rows.Next() {
		err := rows.Scan(
			&reference.UUID,
			&reference.Name,
			&reference.Phone,
			&reference.Relationship,
			&reference.BornDate,
		)

		if err != nil {
			return references, err
		}
		references = append(references, reference)
	}
	return references, nil
}

func (*repoReferences) GetRefPer(ctx context.Context, uuid string) ([]models.References, error) {
	reference := models.References{}
	references := []models.References{}

	query := `
	SELECT  uuid,
			name,
			phone,
			relationship,
			profession,
			company FROM u1ntiesb2kvna45k.references where uuidPerson = ? And isFamiliar = false;`

	rows, err := db.QueryContext(ctx, query, uuid)

	if err != nil {
		return references, err
	}
	for rows.Next() {
		err := rows.Scan(
			&reference.UUID,
			&reference.Name,
			&reference.Phone,
			&reference.Relationship,
			&reference.Profession,
			&reference.Company,
		)

		if err != nil {
			return references, err
		}
		references = append(references, reference)
	}
	return references, nil
}
