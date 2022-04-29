package models

import "github.com/Mynor2397/sqlnulls"

// Person implenta el modelo de la base de datos
type Person struct {
	UUID           sqlnulls.NullString `json:"uuid,omitempty"`
	Fullname       sqlnulls.NullString `json:"fullname,omitempty"`
	CUI            sqlnulls.NullString `json:"cui,omitempty"`
	Partida        sqlnulls.NullString `json:"partida,omitempty"`
	Sueldo         float64             `json:"sueldo,omitempty"`
	AdmissionDate  sqlnulls.NullString `json:"admission_date,omitempty"`
	Job            `json:"job,omitempty"`
	WorkDependency `json:"work_dependency,omitempty"`
	Especiality    `json:"especiality,omitempty"`
	Reubication    `json:"reubication,omitempty"`
	Renglon        sqlnulls.NullString `json:"renglon,omitempty"`
	IsSubstitute   bool                `json:"is_substitute,omitempty"`
	Phone          sqlnulls.NullString `json:"phone,omitempty"`
	DPI            sqlnulls.NullString `json:"dpi,omitempty"`
	NIT            sqlnulls.NullString `json:"nit,omitempty"`
	BornDate       sqlnulls.NullString `json:"born_date,omitempty"`
	Email          sqlnulls.NullString `json:"email,omitempty"`
	Active         bool                `json:"active,omitempty"`
}
