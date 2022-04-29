package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

type repoRequestVacation struct{}

func NewRequestVacation() IRequestVacationStorage {
	return &repoRequestVacation{}
}

type IRequestVacationStorage interface {
	Create(ctx context.Context, requestVacation models.RequestVacation) (models.RequestVacation, error)
	GetRequestsVacations(ctx context.Context, query string, argsQuery []interface{}) ([]models.RequestVacation, error)
	GetOneRequestVacation(ctx context.Context, uuid string) (models.RequestVacation, error)
	UpdateRequestVacation(ctx context.Context, request models.RequestVacation, uuid string) (string, error)
	DeleteRequestVacation(ctx context.Context, uuid string) (string, error)
}

func (*repoRequestVacation) Create(ctx context.Context, request models.RequestVacation) (models.RequestVacation, error) {

	query := "INSERT INTO vacationrequest VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.QueryContext(
		ctx,
		query,
		request.UUIDRequestVacation,
		request.Register,
		request.SubmittedAt,
		request.ModifiedAt,
		request.LastYearVacation,
		request.VacationYearRequest,
		request.LastVacationFrom,
		request.LastVacationTo,
		request.VacationFromDate,
		request.VacationToDate,
		request.HasVacationDay,
		request.DaysQuantity,
		request.Observations,
		request.Person.UUID,
		request.PersonServer.UUID,
		request.UUIDUser,
	)

	if err != nil {
		return request, err
	}

	return request, nil
}

func (*repoRequestVacation) GetRequestsVacations(ctx context.Context, query string, argsQuery []interface{}) ([]models.RequestVacation, error) {
	request := models.RequestVacation{}
	requests := []models.RequestVacation{}

	rows, err := db.QueryContext(ctx, query, argsQuery...)
	if err != nil {
		return requests, err
	}

	for rows.Next() {
		err := rows.Scan(&request.UUIDRequestVacation, &request.Person.Fullname, &request.Person.CUI, &request.SubmittedAt)
		if err != nil {
			return requests, err
		}

		requests = append(requests, request)
	}

	return requests, nil
}

func (*repoRequestVacation) GetOneRequestVacation(ctx context.Context, uuid string) (models.RequestVacation, error) {
	request := models.RequestVacation{}

	query := `
	SELECT 
		r.uuid as uuid_request,
		r.submittedAt,
		p.fullname, 
		p.cui,
		p.admissionDate, 
		j.name as job, 
		e.name as especiality, 
		p.partida,
		reub.name as reubication,
		w.name as workdependency,
		r.lastYearVacation,
		r.vacationYearRequest,
		r.lastVacationFrom, 
		r.lastVacationTo,
		r.vacationFromDate,
		r.vacationToDate,
		r.hasVacationDay,
		r.daysQuantity,
		r.observations,
		substi.fullname as substitute,
		substi.uuid as uuid_substitute
	FROM vacationrequest r
	INNER JOIN person p ON r.person_uuid = p.uuid
	LEFT JOIN job j ON p.job = j.uuid
	LEFT JOIN job e ON p.especiality = e.uuid
	LEFT JOIN job w ON p.workdependency = w.uuid
	LEFT JOIN job reub ON p.reubication = reub.uuid
	LEFT JOIN person substi ON r.publicServer_uuid = substi.uuid
	WHERE r.uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&request.UUIDRequestVacation,
		&request.SubmittedAt,
		&request.Person.Fullname,
		&request.Person.CUI,
		&request.Person.AdmissionDate,
		&request.Person.Job.Name,
		&request.Person.Especiality.Name,
		&request.Person.Partida,
		&request.Person.Reubication.Name,
		&request.Person.WorkDependency.Name,
		&request.LastYearVacation,
		&request.VacationYearRequest,
		&request.LastVacationFrom,
		&request.LastVacationTo,
		&request.VacationFromDate,
		&request.VacationToDate,
		&request.HasVacationDay,
		&request.DaysQuantity,
		&request.Observations,
		&request.PersonServer.Fullname,
		&request.PersonServer.UUID,
	)

	if err == sql.ErrNoRows {
		return request, lib.ErrSQL404
	}
	if err != nil {
		return request, err
	}

	return request, nil
}

func (*repoRequestVacation) UpdateRequestVacation(ctx context.Context, request models.RequestVacation, uuid string) (string, error) {

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
		request.ModifiedAt,
		request.LastYearVacation,
		request.VacationYearRequest,
		request.LastVacationFrom,
		request.LastVacationTo,
		request.VacationFromDate,
		request.VacationToDate,
		request.HasVacationDay,
		request.DaysQuantity,
		request.Observations,
		request.PersonServer.UUID,
		uuid,
	)

	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (*repoRequestVacation) DeleteRequestVacation(ctx context.Context, uuid string) (string, error) {
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
