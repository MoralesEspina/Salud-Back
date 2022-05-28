package models

import "github.com/Mynor2397/sqlnulls"

// References implenta el modelo de la base de datos
type References struct {
	UUID         sqlnulls.NullString `json:"uuid,omitempty"`
	UuidPerson   sqlnulls.NullString `json:"uuidperson,omitempty"`
	Name         sqlnulls.NullString `json:"name,omitempty"`
	Phone        sqlnulls.NullString `json:"phone,omitempty"`
	Relationship sqlnulls.NullString `json:"relationship,omitempty"`
	BornDate     sqlnulls.NullString `json:"bornDate,omitempty"`
	Profession   sqlnulls.NullString `json:"profession,omitempty"`
	Company      sqlnulls.NullString `json:"company,omitempty"`
	IsFamiliar   bool                `json:"isFamiliar,omitempty"`
}
