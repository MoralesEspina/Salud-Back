package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/Mynor2397/sqlnulls"
)

// NewPersonStorage  constructor para userStorage
func NewPersonStorage() PersonStorage {
	return &repoPerson{
		limit:  10,
		offset: 1,
	}
}

type repoPerson struct {
	limit  int
	offset int
	// p PersonStorage
}

type PersonStorage interface {
	Create(ctx context.Context, person models.Person) (models.Person, error)
	GetOne(ctx context.Context, uuid string) (models.Person, error)
	GetMany(ctx context.Context, filter string) ([]models.Person, error)
	GetManyWithFullInformation(ctx context.Context, filter string) ([]models.Person, error)

	Update(ctx context.Context, uuid string, person models.Person) (string, error)
	PaginationQuery(page, limit int) *repoPerson

	CreateSubstitute(ctx context.Context, person models.Person) (models.Person, error)
	GetOneSubstitute(ctx context.Context, uuidPerson string) (models.Person, error)
	GetSubstitutes(ctx context.Context) ([]models.Person, error)

	GetNamePerson(ctx context.Context) ([]models.Person, error)
}

func (*repoPerson) Create(ctx context.Context, person models.Person) (models.Person, error) {
	query := `INSERT INTO person VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	_, err := db.QueryContext(
		ctx,
		query,
		person.UUID,
		person.Fullname,
		person.CUI,
		person.Partida,
		person.Sueldo,
		person.AdmissionDate,
		person.Job.UUIDJob,
		person.WorkDependency.UUIDWork,
		person.Especiality.UUIDEspeciality,
		person.Reubication.UUIDReubication,
		person.Renglon,
		false,
		person.Phone,
		person.DPI,
		person.NIT,
		person.BornDate,
		person.Email,
		person.Active,
	)

	if err != nil {
		return person, err
	}

	return person, nil
}

func (*repoPerson) GetOne(ctx context.Context, uuid string) (models.Person, error) {
	person := models.Person{}

	query := `
	SELECT 
    	p.uuid, 
    	p.fullname, 
    	p.cui,
		p.partida,
    	p.sueldo,
    	p.admissionDate,
    	j.name as job,
    	j.uuid as uuid_job,
    	w.name as workdependency,
    	w.uuid as uuid_workdependency,
    	esp.name as especiality,
    	esp.uuid as uuid_especiality,
    	reu.name as reubication,
    	reu.uuid as uuid_reubication,
		p.collective,
		p.phone,
		p.DPI,
		p.NIT,
		p.bornDate,
		p.email,
		p.active
    FROM person p
    	INNER JOIN job w ON p.workdependency = w.uuid
		LEFT JOIN job j ON p.job = j.uuid
    	LEFT JOIN job esp ON p.especiality = esp.uuid
    	LEFT JOIN job reu ON p.reubication = reu.uuid
	WHERE isPublicServer = false AND p.uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&person.UUID,
		&person.Fullname,
		&person.CUI,
		&person.Partida,
		&person.Sueldo,
		&person.AdmissionDate,
		&person.Job.Name,
		&person.Job.UUIDJob,
		&person.WorkDependency.Name,
		&person.WorkDependency.UUIDWork,
		&person.Especiality.Name,
		&person.Especiality.UUIDEspeciality,
		&person.Reubication.Name,
		&person.Reubication.UUIDReubication,
		&person.Renglon,
		&person.Phone,
		&person.DPI,
		&person.NIT,
		&person.BornDate,
		&person.Email,
		&person.Active,
	)

	if err == sql.ErrNoRows {
		return person, lib.ErrNotFound
	}

	if err != nil {
		return person, err
	}

	return person, nil
}

func (p *repoPerson) GetMany(ctx context.Context, filter string) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}

	var query string
	args := []interface{}{}

	if filter != "" {
		filter = "%" + filter + "%"
		query = `SELECT p.uuid, p.fullname, p.cui, p.partida, j.name as job, j.uuid as uuid_job, w.name as work, w.uuid as uuid_work FROM person p
		LEFT JOIN job j ON p.job = j.uuid
        INNER JOIN job w ON p.workdependency = w.uuid
		WHERE fullname LIKE ? AND isPublicServer = false
		ORDER BY p.fullname ASC
		LIMIT ? OFFSET ?;`

		args = append(args, filter, p.limit, p.offset)

	} else {
		query = `SELECT p.uuid, p.fullname, p.cui, p.partida, j.name as job, j.uuid as uuid_job, w.name as work, w.uuid as uuid_work FROM person p
		LEFT JOIN job j ON p.job = j.uuid
        INNER JOIN job w ON p.workdependency = w.uuid
		WHERE isPublicServer = false
		ORDER BY p.fullname ASC
		LIMIT ? OFFSET ?;`
		args = append(args, p.limit, p.offset)
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(&person.UUID, &person.Fullname, &person.CUI, &person.Partida, &person.Job.Name, &person.UUIDJob, &person.WorkDependency.Name, &person.UUIDWork)
		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (*repoPerson) Update(ctx context.Context, uuid string, person models.Person) (string, error) {

	query := `UPDATE person SET 
					fullname = ?, 
					cui = ?, 
					partida = ?, 
					sueldo = ?, 
					admissionDate = ?, 
					job = ?, 
					workdependency = ?, 
					especiality = ?, 
					reubication = ?,
					collective = ?,
					phone = ?,
					DPI = ?,
					NIT = ?,
					bornDate = ?,
					email = ?,
					active = ? `

	query += " WHERE uuid = ?;"

	_, err := db.QueryContext(
		ctx,
		query,
		person.Fullname,
		person.CUI,
		person.Partida,
		person.Sueldo,
		person.AdmissionDate,
		person.Job.UUIDJob,
		person.WorkDependency.UUIDWork,
		person.Especiality.UUIDEspeciality,
		person.Reubication.UUIDReubication,
		person.Renglon,
		person.Phone,
		person.DPI,
		person.NIT,
		person.BornDate,
		person.Email,
		person.Active,
		uuid,
	)

	if err != nil {
		return "", err
	}

	return string(person.UUID), nil
}

// substitute
func (*repoPerson) CreateSubstitute(ctx context.Context, person models.Person) (models.Person, error) {
	query := `INSERT INTO person (uuid, fullname, isPublicServer) VALUES(?,?,?);`

	_, err := db.QueryContext(
		ctx,
		query,
		person.UUID,
		person.Fullname,
		person.IsSubstitute,
	)

	if err != nil {
		return person, err
	}

	return person, nil
}

func (*repoPerson) GetSubstitutes(ctx context.Context) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}
	query := "SELECT uuid, fullname, isPublicServer FROM person WHERE isPublicServer is true;"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(&person.UUID, &person.Fullname, &person.IsSubstitute)
		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}
	return persons, nil
}

func (*repoPerson) GetOneSubstitute(ctx context.Context, uuidPerson string) (models.Person, error) {
	person := models.Person{}

	query := `SELECT uuid, fullname, isPublicServer FROM person WHERE isPublicServer is true AND uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuidPerson).Scan(
		&person.UUID,
		&person.Fullname,
		&person.IsSubstitute,
	)

	if err == sql.ErrNoRows {
		return person, lib.ErrNotFound
	}

	if err != nil {
		return person, err
	}

	return person, nil
}

func ValidateExistePerson(ctx context.Context, uuid string) (sqlnulls.NullString, error) {
	var nameExist sqlnulls.NullString
	var Count int
	err := db.QueryRowContext(ctx, "SELECT COUNT(*), fullname FROM person WHERE uuid = ? GROUP BY fullname;", uuid).Scan(&Count, &nameExist)
	if err == sql.ErrNoRows {
		return "", lib.ErrNotFound
	}

	if err != nil {
		return "", err
	}

	if Count == 0 {
		return "", lib.ErrNotFound
	}

	return nameExist, nil
}

func (p *repoPerson) GetManyWithFullInformation(ctx context.Context, filter string) ([]models.Person, error) {
	var query string
	person := models.Person{}
	persons := []models.Person{}
	args := []interface{}{}
	if filter != "" {
		filter = "%" + filter + "%"
		query = `SELECT
					p.uuid,
					p.cui,
					p.fullname,
					j.name as job,
					e.name as especiality,
					p.partida,
					p.admissionDate,
					r.name as reubication,
					w.name as workdependency,
					p.collective,
					p.phone,
					p.DPI,
					p.NIT,
					p.bornDate,
					p.email,
					p.active
				FROM person p
				LEFT JOIN job j ON p.job = j.uuid
				LEFT JOIN job e ON p.especiality = e.uuid
				INNER JOIN job w ON p.workdependency = w.uuid
				LEFT JOIN job r ON p.reubication = r.uuid
				WHERE 
				p.fullname LIKE ? OR
				p.cui LIKE ? OR
				e.name LIKE ? OR
				j.name LIKE ? OR
				p.partida LIKE ? OR
				r.name LIKE ? OR
				w.name LIKE ?

				AND isPublicServer = false
				LIMIT ? OFFSET ?;`

		args = append(args, filter, filter, filter, filter, filter, filter, filter, p.limit, p.offset)
	} else {
		query = `SELECT
					p.uuid,
					p.cui,
					p.fullname,
					j.name as job,
					e.name as especiality,
					p.partida,
					p.admissionDate,
					r.name as reubication,
					w.name as workdependency,
					p.collective,
					p.phone,
					p.DPI,
					p.NIT,
					p.bornDate,
					p.email,
					p.active
				FROM person p
				LEFT JOIN job j ON p.job = j.uuid
				LEFT JOIN job e ON p.especiality = e.uuid
				INNER JOIN job w ON p.workdependency = w.uuid
				LEFT JOIN job r ON p.reubication = r.uuid
				WHERE isPublicServer = false
				LIMIT ? OFFSET ?;`
		args = append(args, p.limit, p.offset)
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(
			&person.UUID,
			&person.CUI,
			&person.Fullname,
			&person.Job.Name,
			&person.Especiality.Name,
			&person.Partida,
			&person.AdmissionDate,
			&person.Reubication.Name,
			&person.WorkDependency.Name,
			&person.Renglon,
			&person.Phone,
			&person.DPI,
			&person.NIT,
			&person.BornDate,
			&person.Email,
			&person.Active,
		)

		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (p *repoPerson) PaginationQuery(page, limit int) *repoPerson {
	if limit != 0 {
		p.limit = limit
	}

	if page < 1 {
		page = 1
	}

	p.offset = (page - 1) * limit

	return p
}

func (*repoPerson) GetNamePerson(ctx context.Context) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}
	query := "SELECT uuid, fullname FROM person WHERE isPublicServer is false;"

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
