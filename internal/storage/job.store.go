package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

func NewJobStorage() IJobStorage {
	return &repoJob{}
}

type repoJob struct{}

type IJobStorage interface {
	ManyJobs(ctx context.Context) ([]models.Job, error)
	OneJob(ctx context.Context, uuid string) (models.Job, error)
	CreateJob(ctx context.Context, job models.Job) (string, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, job models.Job, uuid string) (models.Job, error)
}

func (*repoJob) CreateJob(ctx context.Context, job models.Job) (string, error) {
	query := "INSERT INTO job (uuid, name, isJob) VALUES (?, ?, true)"

	_, err := db.QueryContext(ctx, query, job.UUIDJob, job.Name)
	if err != nil {
		return "", err
	}

	return string(job.UUIDJob), nil
}

func (*repoJob) ManyJobs(ctx context.Context) ([]models.Job, error) {
	query := "SELECT uuid, name FROM job WHERE isJob = true ORDER BY name ASC;"
	job := models.Job{}
	jobs := []models.Job{}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return jobs, err
	}

	for rows.Next() {
		if err := rows.Scan(&job.UUIDJob, &job.Name); err != nil {
			return jobs, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (*repoJob) OneJob(ctx context.Context, uuid string) (models.Job, error) {
	query := "SELECT uuid, name FROM job WHERE uuid = ?;"
	job := models.Job{}

	err := db.QueryRowContext(ctx, query, uuid).Scan(&job.UUIDJob, &job.Name)
	if err == sql.ErrNoRows {
		return job, lib.ErrNotFound
	}

	if err != nil {
		return job, err
	}

	return job, nil
}

func (*repoJob) Delete(ctx context.Context, uuid string) error {
	query := "DELETE FROM job WHERE uuid = ? AND isJob = true;"

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

func (*repoJob) Update(ctx context.Context, job models.Job, uuid string) (models.Job, error) {
	queryVerify := "SELECT COUNT(name) FROM job WHERE uuid = ?;"
	var isInDB int
	err := db.QueryRowContext(ctx, queryVerify, uuid).Scan(&isInDB)
	if err != nil {
		return job, err
	}

	if isInDB == 0 {
		return job, lib.ErrNotFound
	}

	queryUpdate := "UPDATE job SET name = ? WHERE uuid = ?;"
	_, err = db.ExecContext(ctx, queryUpdate, job.Name, uuid)
	if err != nil {
		return job, err
	}

	job.UUIDJob.Scan(uuid)
	return job, nil
}
