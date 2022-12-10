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
	GetBosssesOne(ctx context.Context) ([]models.Person, error)
	GetBosssesTwo(ctx context.Context) ([]models.Person, error)
	GetPermissionsBossOne(ctx context.Context, uuid string) ([]models.Permission, error)
	GetPermissionsBossTwo(ctx context.Context, uuid string) ([]models.Permission, error)
	GetUserPermissionsActives(ctx context.Context, uuid string) ([]models.Permission, error)
	GetUserPermissions(ctx context.Context, uuid string) ([]models.Permission, error)
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
		request.Status,
		request.Reason,
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

	query := `SELECT uuid, submittedAt, permissionDate, uuidPerson, bossOne, bossTwo, motive, statusBossOne, StatusBossTwo, reason, status FROM permission where uuid = ?`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&request.Uuid,
		&request.SubmittedAt,
		&request.PermissionDate,
		&request.UuidPerson,
		&request.BossOne,
		&request.BossTwo,
		&request.Motive,
		&request.StatusBossOne,
		&request.StatusBossTwo,
		&request.Reason,
		&request.Status,
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
	UPDATE permission
	SET
		modifiedAt = ?,
		statusBossOne = ?,
		StatusBossTwo = ?,
		reason = ?,
		status = ?
		WHERE uuid = ?;
	`
	_, err := db.QueryContext(
		ctx,
		query,
		request.ModifiedAt,
		request.StatusBossOne,
		request.StatusBossTwo,
		request.Reason,
		request.Status,
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

func (*repoPermission) GetBosssesOne(ctx context.Context) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}
	query := `	SELECT u.uuid, p.fullname FROM person p
				JOIN user u WHERE p.uuid = u.uuidPerson
				AND u.rol_id IN (4,6);`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(&person.UUID, &person.Fullname)
		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}
	return persons, nil
}

func (*repoPermission) GetBosssesTwo(ctx context.Context) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}
	query := `SELECT u.uuid, p.fullname FROM person p
				 JOIN user u WHERE p.uuid = u.uuidPerson
				 AND u.rol_id = 6;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(&person.UUID, &person.Fullname)
		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}
	return persons, nil
}

func (*repoPermission) GetPermissionsBossOne(ctx context.Context, uuid string) ([]models.Permission, error) {
	permission := models.Permission{}
	permissions := []models.Permission{}
	query := `	SELECT r.uuid, r.submittedAt, r.permissionDate, p.fullname FROM permission r INNER JOIN person p ON r.uuidPerson = p.uuid 
				WHERE r.bossOne = (Select uuid from user u where u.uuidPerson = ?) 
				AND r.statusBossOne LIKE 'En Espera'
				ORDER BY r.submittedAt ASC;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		err := rows.Scan(&permission.Uuid, &permission.SubmittedAt, &permission.PermissionDate, &permission.Person.Fullname)
		if err != nil {
			return permissions, err
		}

		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (*repoPermission) GetPermissionsBossTwo(ctx context.Context, uuid string) ([]models.Permission, error) {
	permission := models.Permission{}
	permissions := []models.Permission{}
	query := `	SELECT r.uuid, r.submittedAt, p.fullname as applicant, r.permissionDate, pe.fullname as bossOne from permission r JOIN person p ON r.uuidPerson = p.uuid JOIN user u ON r.bossOne = u.uuid JOIN person pe ON u.uuidPerson = pe.uuid
				WHERE r.bossTwo = (Select uuid from user u where u.uuidPerson = ?) 
				AND r.statusBossOne LIKE 'Aceptada'
				AND r.StatusBossTwo LIKE 'En Espera'
				ORDER BY r.submittedAt ASC;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		err := rows.Scan(&permission.Uuid, &permission.SubmittedAt, &permission.Person.Fullname, &permission.PermissionDate, &permission.BossOne)
		if err != nil {
			return permissions, err
		}

		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (*repoPermission) GetUserPermissionsActives(ctx context.Context, uuid string) ([]models.Permission, error) {
	permission := models.Permission{}
	permissions := []models.Permission{}
	query := `	SELECT r.uuid, r.submittedAt, r.permissionDate, r.statusBossOne, r.StatusBossTwo, r.status, p.fullname as bossOne, pe.fullname as bossTwo FROM permission r JOIN person p ON r.uuidPerson = p.uuid JOIN user u ON r.bossOne = u.uuid JOIN person pe ON u.uuidPerson = pe.uuid JOIN user us ON r.bossTwo = us.uuid
				WHERE r.uuidPerson = ?
				AND r.status LIKE 'En Espera'
				ORDER BY r.submittedAt ASC;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		err := rows.Scan(&permission.Uuid, &permission.SubmittedAt, &permission.PermissionDate, &permission.StatusBossOne, &permission.StatusBossTwo, &permission.Status, &permission.BossOne, &permission.BossTwo)
		if err != nil {
			return permissions, err
		}

		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (*repoPermission) GetUserPermissions(ctx context.Context, uuid string) ([]models.Permission, error) {
	permission := models.Permission{}
	permissions := []models.Permission{}
	query := `	SELECT r.uuid, r.submittedAt, r.permissionDate, r.status, p.fullname as bossOne, pe.fullname as bossTwo FROM permission r JOIN person p ON r.uuidPerson = p.uuid JOIN user u ON r.bossOne = u.uuid JOIN person pe ON u.uuidPerson = pe.uuid JOIN user us ON r.bossTwo = us.uuid
				WHERE r.uuidPerson = ?
				AND r.status NOT LIKE 'En Espera'
				ORDER BY r.submittedAt ASC;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		err := rows.Scan(&permission.Uuid, &permission.SubmittedAt, &permission.PermissionDate, &permission.Status, &permission.BossOne, &permission.BossTwo)
		if err != nil {
			return permissions, err
		}

		permissions = append(permissions, permission)
	}
	return permissions, nil
}
