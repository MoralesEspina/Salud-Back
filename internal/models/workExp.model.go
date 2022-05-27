package models

import "github.com/Mynor2397/sqlnulls"

// WorkExp implenta el modelo de la base de datos
type WorkExp struct {
	UUID       sqlnulls.NullString `json:"uuid,omitempty"`
	UuidPerson sqlnulls.NullString `json:"uuidperson,omitempty"`
	Direction  sqlnulls.NullString `json:"direction,omitempty"`
	Phone      sqlnulls.NullString `json:"phone,omitempty"`
	Reason     sqlnulls.NullString `json:"reason,omitempty"`
	DateOf     sqlnulls.NullString `json:"dateof,omitempty"`
	DateTo     sqlnulls.NullString `json:"dateto,omitempty"`
	Job        sqlnulls.NullString `json:"job,omitempty"`
	BossName   sqlnulls.NullString `json:"bossname,omitempty"`
	Sector     sqlnulls.NullString `json:"sector,omitempty"`
	Salary     float64             `json:"salary,omitempty"`
	WorkExpCol sqlnulls.NullString `json:"workexpcol,omitempty"`
}
