package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

// NewCurriculumStorage  constructor para CurriculumStorage
func NewCurriculumStorage() CurriculumStorage {
	return &repoCurriculum{}
}

type repoCurriculum struct {
}

type CurriculumStorage interface {
	Create(ctx context.Context, curriculum models.Curriculum) (models.Curriculum, error)
	GetOne(ctx context.Context, uuid string) (models.Curriculum, error)
	Update(ctx context.Context, uuid string, curriculum models.Curriculum) (string, error)
}

func (*repoCurriculum) Create(ctx context.Context, curriculum models.Curriculum) (models.Curriculum, error) {
	query := `INSERT INTO curriculum VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	_, err := db.QueryContext(
		ctx,
		query,
		curriculum.UUID,
		curriculum.UuidPerson,
		curriculum.Direction,
		curriculum.Country,
		curriculum.HomePhone,
		curriculum.BornPlace,
		curriculum.Nacionality,
		curriculum.Municipality,
		curriculum.Village,
		curriculum.WorkPhone,
		curriculum.Age,
		curriculum.CivilStatus,
		curriculum.Etnia,
		curriculum.Passport,
		curriculum.License,
		curriculum.Department,
		curriculum.IGSS,
	)

	if err != nil {
		return curriculum, err
	}

	return curriculum, nil
}

func (*repoCurriculum) GetOne(ctx context.Context, uuid string) (models.Curriculum, error) {
	curriculum := models.Curriculum{}

	query := `
	SELECT 
		c.uuid,
		c.uuidPerson,
        c.direction, 
        c.country,
        c.homePhone,
        c.bornPlace,
        c.nacionality,
        c.municipality, 
		c.village, 
		c.workPhone  ,
		c.age ,
		c.civilStatus ,
		c.etnia ,
		c.passport ,
		c.license ,
		c.department ,
		c.igss,
        p.phone as phone,
		p.DPI as DPI,
        p.NIT as NIT,
        p.bornDate as bornDate,
        p.email as email,
        p.fullname as fullname
    FROM curriculum c
        INNER JOIN person p ON c.uuidPerson = p.uuid
    WHERE p.uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&curriculum.UUID,
		&curriculum.UuidPerson,
		&curriculum.Direction,
		&curriculum.Country,
		&curriculum.HomePhone,
		&curriculum.BornPlace,
		&curriculum.Nacionality,
		&curriculum.Municipality,
		&curriculum.Village,
		&curriculum.WorkPhone,
		&curriculum.Age,
		&curriculum.CivilStatus,
		&curriculum.Etnia,
		&curriculum.Passport,
		&curriculum.License,
		&curriculum.Department,
		&curriculum.IGSS,
		&curriculum.Person.Phone,
		&curriculum.Person.DPI,
		&curriculum.Person.NIT,
		&curriculum.Person.BornDate,
		&curriculum.Person.Email,
		&curriculum.Person.Fullname,
	)

	if err == sql.ErrNoRows {
		return curriculum, lib.ErrNotFound
	}

	if err != nil {
		return curriculum, err
	}

	return curriculum, nil
}

func (*repoCurriculum) Update(ctx context.Context, uuid string, curriculum models.Curriculum) (string, error) {

	query := `UPDATE curriculum SET 
					direction = ?, 
					country = ?, 
					homePhone = ?, 
					bornPlace = ?, 
					nacionality = ?, 
					municipality = ?, 
					village = ?, 
					workPhone = ?, 
					age = ?,
					civilStatus = ?,
					etnia = ?,
					passport = ?,
					license = ? `

	query += " WHERE uuidPerson = ?;"

	_, err := db.QueryContext(
		ctx,
		query,
		curriculum.Direction,
		curriculum.Country,
		curriculum.HomePhone,
		curriculum.BornPlace,
		curriculum.Nacionality,
		curriculum.Municipality,
		curriculum.Village,
		curriculum.WorkPhone,
		curriculum.Age,
		curriculum.CivilStatus,
		curriculum.Etnia,
		curriculum.Passport,
		curriculum.License,
		uuid,
	)

	if err != nil {
		return "", err
	}

	return string(curriculum.UUID), nil
}
