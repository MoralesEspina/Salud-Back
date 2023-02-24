package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

type repoAuthorization struct{}

// NewPersonStorage  constructor para userStorage
func NewAuthorizationStorage() AuthorizationStorage {
	return &repoAuthorization{}
}

type AuthorizationStorage interface {
	Create(ctx context.Context, authorization models.Authorization) (models.Authorization, error)
	GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error)
	GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error)
	UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error)
	GetBosses(ctx context.Context) ([]models.Boss, error)
	UpdateBoss(ctx context.Context, authorization models.Boss, id string) (models.Boss, error)

	// Reportes
	VacationsReport(ctx context.Context, startDateReport, endDateReport string) ([]models.Authorization, error)
}

func (*repoAuthorization) GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error) {
	autorization := models.Authorization{}
	autorizations := []models.Authorization{}

	query := `SELECT a.uuid, a.register, a.submittedAt, p.fullname, p.cui FROM autorization a 
			  INNER JOIN person p ON a.person_idperson = p.uuid
			  ORDER BY submittedAt DESC, register DESC;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return autorizations, err
	}

	for rows.Next() {
		if err := rows.Scan(&autorization.UUIDAuthorization, &autorization.Register, &autorization.SubmittedAt, &autorization.Fullname, &autorization.CUI); err != nil {
			return autorizations, err
		}

		autorizations = append(autorizations, autorization)
	}

	return autorizations, nil
}

func (*repoAuthorization) GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error) {
	autorization := models.Authorization{}

	query := `
	SELECT 
		a.uuid,
		a.person_idperson,
		w.uuid,
		a.register,
		p.cui,
		p.fullname,
		a.submittedAt,
		a.startdate,
		a.enddate,
		a.resumework,
		a.holidays,
		a.totaldays,
		a.pendingdays,
		a.observation,
		a.authorizationyear,
		p.partida,
		w.name as workdependency,
		j.name as job,
		a.personnelOfficer,
		a.personnelOfficerPosition,
		a.personnelOfficerArea,
		a.executiveDirector,
		a.executiveDirectorPosition,
		a.executiveDirectorArea
	FROM autorization a
	INNER JOIN person p ON a.person_idperson = p.uuid
	LEFT JOIN job j ON p.job = j.uuid
	INNER JOIN job w ON p.workdependency = w.uuid
	LEFT JOIN job es ON p.especiality = es.uuid
	LEFT JOIN job reu ON p.reubication = reu.uuid
	WHERE a.uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&autorization.UUIDAuthorization,
		&autorization.UUID,
		&autorization.WorkdependencyUUID,
		&autorization.Register,
		&autorization.CUI,
		&autorization.Fullname,
		&autorization.SubmittedAt,
		&autorization.Startdate,
		&autorization.Enddate,
		&autorization.Resumework,
		&autorization.Holidays,
		&autorization.TotalDays,
		&autorization.Pendingdays,
		&autorization.Observation,
		&autorization.Authorizationyear,
		&autorization.Partida,
		&autorization.WorkDependency.Name,
		&autorization.Job.Name,
		&autorization.PersonnelOfficer,
		&autorization.PersonnelOfficerPosition,
		&autorization.PersonnelOfficerArea,
		&autorization.ExecutiveDirector,
		&autorization.ExecutiveDirectorPosition,
		&autorization.ExecutiveDirectorArea,
	)

	if err == sql.ErrNoRows {
		return autorization, lib.ErrNotFound
	}
	if err != nil {
		return autorization, err
	}

	return autorization, nil
}

func (*repoAuthorization) Create(ctx context.Context, authorization models.Authorization) (models.Authorization, error) {
	authoriza := models.Authorization{}

	query := `
	INSERT INTO
    autorization (
        uuid,
        register,
        submittedAt,
		modifiedAt,
        startdate,
        enddate,
		resumework,
		holidays,
        totaldays,
        pendingdays,
        observation,
        authorizationyear,
        person_idperson,
		user_uuid,
		
		personnelOfficer,
		personnelOfficerPosition,
		personnelOfficerArea,
		executiveDirector,
		executiveDirectorPosition,
		executiveDirectorArea
    )
    VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	trans, err := db.BeginTx(ctx, nil)

	if err != nil {
		return authoriza, err
	}
	defer trans.Rollback()

	_, err = db.QueryContext(
		ctx,
		query,
		authorization.UUIDAuthorization,
		authorization.Register,
		authorization.SubmittedAt,
		authorization.ModifiedAt,
		authorization.Startdate,
		authorization.Enddate,
		authorization.Resumework,
		authorization.Holidays,
		authorization.TotalDays,
		authorization.Pendingdays,
		authorization.Observation,
		authorization.Authorizationyear,
		authorization.UUID,
		authorization.User,

		authorization.PersonnelOfficer,
		authorization.PersonnelOfficerPosition,
		authorization.PersonnelOfficerArea,
		authorization.ExecutiveDirector,
		authorization.ExecutiveDirectorPosition,
		authorization.ExecutiveDirectorArea,
	)
	if err != nil {
		return authoriza, err
	}

	authoriza, err = DataPDFAuthorization(ctx, authorization.UUIDAuthorization, db)
	if err != nil {
		return authoriza, err
	}

	if errtrans := trans.Commit(); errtrans != nil {
		return authoriza, errtrans
	}

	return authoriza, nil
}

func (*repoAuthorization) UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error) {
	authoriza := models.Authorization{}
	trans, err := db.BeginTx(ctx, nil)

	if err != nil {
		return authoriza, err
	}
	defer trans.Rollback()

	query := `
	UPDATE autorization 
	SET submittedAt = ?,
		startdate = ?,
        enddate = ?,
        resumework = ?,
        holidays = ?,
        totaldays = ?,
        pendingdays = ?,
        observation = ?,
        authorizationyear = ?,
        person_idperson = ?,
        personnelOfficer = ?,
        personnelOfficerPosition = ?,
        personnelOfficerArea = ?,
        executiveDirector = ?,
        executiveDirectorPosition = ?,
        executiveDirectorArea = ?
    WHERE uuid = ?`

	_, err = db.QueryContext(ctx, query,
		authorization.SubmittedAt,
		authorization.Startdate,
		authorization.Enddate,
		authorization.Resumework,
		authorization.Holidays,
		authorization.TotalDays,
		authorization.Pendingdays,
		authorization.Observation,
		authorization.Authorizationyear,
		authorization.Person.UUID,

		authorization.PersonnelOfficer,
		authorization.PersonnelOfficerPosition,
		authorization.PersonnelOfficerArea,
		authorization.ExecutiveDirector,
		authorization.ExecutiveDirectorPosition,
		authorization.ExecutiveDirectorArea,
		uuid,
	)
	if err != nil {
		return authoriza, err
	}

	authoriza, err = DataPDFAuthorization(ctx, uuid, db)

	if err != nil {
		return authoriza, err
	}

	if errtrans := trans.Commit(); errtrans != nil {
		return authoriza, errtrans
	}

	return authoriza, nil
}

func DataPDFAuthorization(ctx context.Context, UUIDAuthorization string, trans *sql.DB) (models.Authorization, error) {
	authoriza := models.Authorization{}
	querySelect := `
	SELECT
    	a.register,
    	a.submittedAt,
    	a.startdate,
    	a.enddate,
		a.resumework,
		a.holidays,
    	a.totaldays,
    	a.pendingdays,
    	a.observation,
    	a.authorizationyear,
    	p.partida,
    	w.name as workdependency,
    	p.fullname,
    	p.cui,
		j.name as job,
		
		a.personnelOfficer,
		a.personnelOfficerPosition,
		a.personnelOfficerArea,
		a.executiveDirector,
		a.executiveDirectorPosition,
		a.executiveDirectorArea
	FROM
    	autorization a
    	INNER JOIN person p ON a.person_idperson = p.uuid
		INNER JOIN job j ON p.job = j.uuid
		LEFT JOIN job w ON p.workdependency = w.uuid
    	WHERE a.uuid = ?
	`
	err := db.QueryRowContext(ctx, querySelect, UUIDAuthorization).Scan(
		&authoriza.Register,
		&authoriza.SubmittedAt,
		&authoriza.Startdate,
		&authoriza.Enddate,
		&authoriza.Resumework,
		&authoriza.Holidays,
		&authoriza.TotalDays,
		&authoriza.Pendingdays,
		&authoriza.Observation,
		&authoriza.Authorizationyear,
		&authoriza.Person.Partida,
		&authoriza.WorkDependency.Name,
		&authoriza.Fullname,
		&authoriza.CUI,
		&authoriza.Job.Name,

		&authoriza.PersonnelOfficer,
		&authoriza.PersonnelOfficerPosition,
		&authoriza.PersonnelOfficerArea,
		&authoriza.ExecutiveDirector,
		&authoriza.ExecutiveDirectorPosition,
		&authoriza.ExecutiveDirectorArea,
	)

	if err == sql.ErrNoRows {
		return authoriza, lib.ErrNotFound
	}

	if err != nil {
		return authoriza, err
	}

	return authoriza, nil
}

/*
	FUNCTIONS TO GENERATE REPORTS
*/

func (*repoAuthorization) VacationsReport(ctx context.Context, startDateReport, endDateReport string) ([]models.Authorization, error) {
	authorization := models.Authorization{}
	authorizations := []models.Authorization{}

	query := `
		SELECT p.fullname, w.name, aut.startdate, aut.enddate, aut.resumework FROM autorization aut
		INNER JOIN person p ON aut.person_idperson = p.uuid
		INNER JOIN job w ON p.workdependency = w.uuid
		WHERE startdate >= ? AND enddate <= ?;`

	rows, err := db.QueryContext(ctx, query, startDateReport, endDateReport)
	if err == sql.ErrNoRows {
		return authorizations, lib.ErrNotFound
	}

	if err != nil {
		return authorizations, err
	}

	for rows.Next() {
		if err := rows.Scan(&authorization.Person.Fullname, &authorization.Person.WorkDependency.Name, &authorization.Startdate, &authorization.Enddate, &authorization.Resumework); err != nil {
			return authorizations, err
		}

		authorizations = append(authorizations, authorization)
	}

	return authorizations, nil
}

func (*repoAuthorization) GetBosses(ctx context.Context) ([]models.Boss, error) {
	boss := models.Boss{}
	bosses := []models.Boss{}

	query := `SELECT * FROM bossauthorization;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return bosses, err
	}

	for rows.Next() {
		if err := rows.Scan(&boss.ID, &boss.Name); err != nil {
			return bosses, err
		}

		bosses = append(bosses, boss)
	}

	return bosses, nil
}

func (*repoAuthorization) UpdateBoss(ctx context.Context, bosses models.Boss, id string) (models.Boss, error) {
	boss := models.Boss{}
	trans, err := db.BeginTx(ctx, nil)

	if err != nil {
		return boss, err
	}
	defer trans.Rollback()

	query := `
	UPDATE bossauthorization 
	SET name = ?
    WHERE id = ?`

	_, err = db.QueryContext(ctx, query,
		bosses.Name,
		id,
	)
	if err != nil {
		return boss, err
	}

	return boss, nil
}
