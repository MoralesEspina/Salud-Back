package models

import "github.com/Mynor2397/sqlnulls"

// Curriculum implenta el modelo de la base de datos
type Curriculum struct {
	UUID         sqlnulls.NullString `json:"uuid,omitempty"`
	UuidPerson   sqlnulls.NullString `json:"uuidPerson,omitempty"`
	Direction    sqlnulls.NullString `json:"direction,omitempty"`
	Country      sqlnulls.NullString `json:"country,omitempty"`
	HomePhone    sqlnulls.NullString `json:"homephone,omitempty"`
	BornPlace    sqlnulls.NullString `json:"bornPlace,omitempty"`
	Nacionality  sqlnulls.NullString `json:"nacionality,omitempty"`
	Municipality sqlnulls.NullString `json:"municipality,omitempty"`
	Village      sqlnulls.NullString `json:"village,omitempty"`
	WorkPhone    sqlnulls.NullString `json:"workPhone,omitempty"`
	Age          sqlnulls.NullString `json:"age,omitempty"`
	CivilStatus  sqlnulls.NullString `json:"civilStatus,omitempty"`
	Etnia        sqlnulls.NullString `json:"etnia,omitempty"`
	Passport     sqlnulls.NullString `json:"passport,omitempty"`
	License      sqlnulls.NullString `json:"license,omitempty"`
	Department   sqlnulls.NullString `json:"department,omitempty"`
	IGSS         sqlnulls.NullString `json:"iggs,omitempty"`
	Person
}
