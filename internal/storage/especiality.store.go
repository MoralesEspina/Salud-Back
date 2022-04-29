package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

func NewEspecialityStorage() IEspecialityStorage {
	return &repoEspeciality{}
}

type repoEspeciality struct{}

type IEspecialityStorage interface {
	Especialities(ctx context.Context) ([]models.Especiality, error)
	OneEspeciality(ctx context.Context, uuid string) (models.Especiality, error)
	CreateEspeciality(ctx context.Context, especiality models.Especiality) (string, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, especiality models.Especiality, uuid string) (models.Especiality, error)
}

func (*repoEspeciality) CreateEspeciality(ctx context.Context, especiality models.Especiality) (string, error) {
	query := "INSERT INTO job (uuid, name, isEspeciality) VALUES (?, ?, true);"

	_, err := db.QueryContext(ctx, query, especiality.UUIDEspeciality, especiality.Name)
	if err != nil {
		return "", err
	}

	return string(especiality.UUIDEspeciality), nil
}

func (*repoEspeciality) Especialities(ctx context.Context) ([]models.Especiality, error) {
	query := "SELECT uuid, name FROM job WHERE isEspeciality = true ORDER BY name ASC;"
	especiality := models.Especiality{}
	especialities := []models.Especiality{}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return especialities, err
	}

	for rows.Next() {
		if err := rows.Scan(&especiality.UUIDEspeciality, &especiality.Name); err != nil {
			return especialities, err
		}

		especialities = append(especialities, especiality)
	}

	return especialities, nil
}

func (*repoEspeciality) OneEspeciality(ctx context.Context, uuid string) (models.Especiality, error) {
	query := "SELECT uuid, name FROM job WHERE uuid = ?;"
	especiality := models.Especiality{}

	err := db.QueryRowContext(ctx, query, uuid).Scan(&especiality.UUIDEspeciality, &especiality.Name)
	if err == sql.ErrNoRows {
		return especiality, lib.ErrNotFound
	}

	if err != nil {
		return especiality, err
	}

	return especiality, nil
}

func (*repoEspeciality) Delete(ctx context.Context, uuid string) error {
	query := "DELETE FROM job WHERE uuid = ? AND isEspeciality = true;"

	rows, err := db.ExecContext(ctx, query, uuid)
	if err != nil {
		return lib.ExtractMysqlError(err)
	}

	resultDelete, _ := rows.RowsAffected()
	if resultDelete == 0 {
		return lib.ErrNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (*repoEspeciality) Update(ctx context.Context, especiality models.Especiality, uuid string) (models.Especiality, error) {
	queryVerify := "SELECT COUNT(*) FROM job WHERE uuid = ?;"
	var isInDB int
	err := db.QueryRowContext(ctx, queryVerify, uuid).Scan(&isInDB)
	if err != nil {
		return especiality, err
	}

	if isInDB == 0 {
		return especiality, lib.ErrNotFound
	}

	queryUpdate := "UPDATE job SET name = ? WHERE uuid = ?;"
	_, err = db.ExecContext(ctx, queryUpdate, especiality.Name, uuid)
	if err != nil {
		return especiality, err
	}

	especiality.UUIDEspeciality.Scan(uuid)
	return especiality, nil
}
