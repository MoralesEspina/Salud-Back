package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

type repoPermission struct{}

func NewPermission() IPermissionStorage {
	return &repoPermission{}
}

type IPermissionStorage interface {
	Create(ctx context.Context, Permission models.Permission) (models.Permission, error)
	GetPermissions(ctx context.Context, query string, argsQuery []interface{}) ([]models.Permission, error)
	GetOnePermission(ctx context.Context, uuid string) (models.Permission, error)
	UpdatePermission(ctx context.Context, request models.Permission, uuid string) (string, error)
	DeletePermission(ctx context.Context, uuid string) (string, error)
}

func (*repoPermission) Create(ctx context.Context, request models.Permission) (models.Permission, error) {

	query := "INSERT INTO permission VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.QueryContext(
		ctx,
		query,
		request.Uuid,
		request.Register,
		request.SubmittedAt,
		request.ModifiedAt,
		request.PermissionDate,
		request.UuidPerson,
		request.BossOne,
		request.BossTwo,
		request.Motive,
		request.StatusBossOne,
		request.StatusBossTwo,
		request.Reason,
		request.Status,
	)

	if err != nil {
		return request, err
	}

	return request, nil
}

func (*repoPermission) GetPermissions(ctx context.Context, query string, argsQuery []interface{}) ([]models.Permission, error) {
	request := models.Permission{}
	requests := []models.Permission{}

	rows, err := db.QueryContext(ctx, query, argsQuery...)
	if err != nil {
		return requests, err
	}

	for rows.Next() {
		err := rows.Scan(&request.Uuid, &request.SubmittedAt, &request.UuidPerson, &request.Reason)
		if err != nil {
			return requests, err
		}

		requests = append(requests, request)
	}

	return requests, nil
}

func (*repoPermission) GetOnePermission(ctx context.Context, uuid string) (models.Permission, error) {
	request := models.Permission{}

	query := `
	SELECT 
		pe.uuid as uuid_request,
		pe.submmittedAt,
		pe.permissionDate, 
		p.uuidPerson,
		pe.bossOne, 
		pe.bossTwo, 
		pe.reasson, 
		pe.document, 
		pe.statusBossOne, 
		pe.statusBossOne, 
		pe.status, 
		WHERE pe.uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&request.Uuid,
		&request.SubmittedAt,
	)

	if err == sql.ErrNoRows {
		return request, lib.ErrSQL404
	}
	if err != nil {
		return request, err
	}

	return request, nil
}

func (*repoPermission) UpdatePermission(ctx context.Context, request models.Permission, uuid string) (string, error) {

	query := `
	UPDATE vacationrequest
	SET
		modifiedAt = ?,
		lastYearVacation = ?,
		vacationYearRequest = ?,
		lastVacationFrom = ?,
		lastVacationTo = ?,
		vacationFromDate = ?,
		vacationToDate = ?,
		hasVacationDay = ?,
		daysQuantity = ?,
		observations = ?,
		publicServer_uuid = ?
	WHERE uuid = ?;
	`
	_, err := db.QueryContext(
		ctx,
		query,

		uuid,
	)

	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (*repoPermission) DeletePermission(ctx context.Context, uuid string) (string, error) {
	queryUpdate := "DELETE FROM vacationrequest WHERE uuid = ?;"

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