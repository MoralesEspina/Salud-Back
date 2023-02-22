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
	GetPermissions(ctx context.Context, startDate, endDate, status string) ([]models.Permission, error)
	GetOnePermission(ctx context.Context, uuid string) (models.Permission, error)
	GetOnePermissionWithName(ctx context.Context, uuid string) (models.Permission, error)
	UpdatePermission(ctx context.Context, request models.Permission, uuid, rol string) (string, error)
	DeletePermission(ctx context.Context, uuid string) (string, error)
	GetBosssesOne(ctx context.Context) ([]models.Person, error)
	GetBosssesTwo(ctx context.Context) ([]models.Person, error)
	GetPermissionsBossOne(ctx context.Context, uuid string) ([]models.Permission, error)
	GetPermissionsBossTwo(ctx context.Context, uuid string) ([]models.Permission, error)
	GetUserPermissionsActives(ctx context.Context, uuid string) ([]models.Permission, error)
	GetUserPermissions(ctx context.Context, uuid string) ([]models.Permission, error)
}

func (*repoPermission) Create(ctx context.Context, request models.Permission) (models.Permission, error) {
	query := "INSERT INTO permission VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

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
		request.Document,
	)

	if err != nil {
		return request, err
	}

	return request, nil
}

func (*repoPermission) GetPermissions(ctx context.Context, startDate, endDate, status string) ([]models.Permission, error) {
	permission := models.Permission{}
	permissions := []models.Permission{}
	query := `	SELECT r.uuid, r.submittedAt, r.permissionDate, p.fullname, r.status, r.uuidPerson FROM permission r
				INNER JOIN person p ON r.uuidPerson = p.uuid
				WHERE r.submittedAt >= ? AND r.submittedAt <= ? AND r.status = ?
				ORDER BY r.submittedAt DESC`

	rows, err := db.QueryContext(ctx, query, startDate, endDate, status)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		err := rows.Scan(&permission.Uuid, &permission.SubmittedAt, &permission.PermissionDate, &permission.Fullname, &permission.Status, &permission.UuidPerson)
		if err != nil {
			return permissions, err
		}

		permissions = append(permissions, permission)
	}
	return permissions, nil

}

func (*repoPermission) GetOnePermission(ctx context.Context, uuid string) (models.Permission, error) {
	request := models.Permission{}

	query := `SELECT p.uuid, p.submittedAt, p.permissionDate, p.uuidPerson, p.bossOne, p.bossTwo, p.motive, p.statusBossOne, p.StatusBossTwo, p.status, p.reason, p.document FROM permission as p where uuid = ?`

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
		&request.Status,
		&request.Reason,
		&request.Document,
	)

	if err == sql.ErrNoRows {
		return request, lib.ErrSQL404
	}
	if err != nil {
		return request, err
	}

	return request, nil
}

func (*repoPermission) GetOnePermissionWithName(ctx context.Context, uuid string) (models.Permission, error) {
	request := models.Permission{}

	query := `SELECT p.uuid, p.submittedAt, p.permissionDate, p.uuidPerson, r.fullname as bossOne, re.fullname as bossTwo, p.motive, p.statusBossOne, 
				p.StatusBossTwo, p.status, p.reason, p.document FROM permission as p JOIN user u ON p.bossOne = u.uuid JOIN person r ON u.uuidPerson = r.uuid 
				JOIN user us ON p.bossTwo = us.uuid JOIN person re ON us.uuidPerson = re.uuid 
				where p.uuid = ?`

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
		&request.Status,
		&request.Reason,
		&request.Document,
	)

	if err == sql.ErrNoRows {
		return request, lib.ErrSQL404
	}
	if err != nil {
		return request, err
	}

	return request, nil
}

func (*repoPermission) UpdatePermission(ctx context.Context, request models.Permission, uuid, rol string) (string, error) {
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
	query := `	SELECT r.uuid, r.submittedAt, r.permissionDate, r.uuidPerson, r.statusBossOne, r.StatusBossTwo, r.status, p.fullname as bossOne, pe.fullname as bossTwo FROM permission r JOIN user u ON r.bossOne = u.uuid JOIN person p ON u.uuidPerson = p.uuid JOIN user us ON r.bossTwo = us.uuid JOIN person pe ON us.uuidPerson = pe.uuid 
				WHERE r.uuidPerson = ?
				AND r.status LIKE 'En Espera'
				ORDER BY r.submittedAt ASC;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		err := rows.Scan(&permission.Uuid, &permission.SubmittedAt, &permission.PermissionDate, &permission.UuidPerson, &permission.StatusBossOne, &permission.StatusBossTwo, &permission.Status, &permission.BossOne, &permission.BossTwo)
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
	query := `	SELECT r.uuid, r.submittedAt, r.permissionDate, r.uuidPerson, r.status, p.fullname as bossOne, pe.fullname as bossTwo FROM permission r JOIN user u ON r.bossOne = u.uuid JOIN person p ON u.uuidPerson = p.uuid JOIN user us ON r.bossTwo = us.uuid JOIN person pe ON us.uuidPerson = pe.uuid 
				WHERE r.uuidPerson = ?
				AND r.status NOT LIKE 'En Espera'
				ORDER BY r.submittedAt DESC;`

	rows, err := db.QueryContext(ctx, query, uuid)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		err := rows.Scan(&permission.Uuid, &permission.SubmittedAt, &permission.PermissionDate, &permission.UuidPerson, &permission.Status, &permission.BossOne, &permission.BossTwo)
		if err != nil {
			return permissions, err
		}

		permissions = append(permissions, permission)
	}
	return permissions, nil
}
