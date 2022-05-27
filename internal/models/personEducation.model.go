package models

import "github.com/Mynor2397/sqlnulls"

// PersonEducation implenta el modelo de la base de datos
type PersonEducation struct {
	UUID          sqlnulls.NullString `json:"uuid,omitempty"`
	UuidPerson    sqlnulls.NullString `json:"uuidperson,omitempty"`
	Country       sqlnulls.NullString `json:"country,omitempty"`
	Establishment sqlnulls.NullString `json:"establishment,omitempty"`
	PeriodOf      sqlnulls.NullString `json:"periodof,omitempty"`
	PeriodTo      sqlnulls.NullString `json:"periodto,omitempty"`
	Certificate   sqlnulls.NullString `json:"certificate,omitempty"`
	Status        sqlnulls.NullString `json:"status,omitempty"`
	Grade         sqlnulls.NullString `json:"grade,omitempty"`
}
