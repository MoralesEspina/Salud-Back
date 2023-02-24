package models

import "github.com/Mynor2397/sqlnulls"

// Person implenta el modelo de la base de datos
type Boss struct {
	ID           sqlnulls.NullString `json:"id,omitempty"`
	NameBoss     sqlnulls.NullString `json:"nameboss,omitempty"`
	NameDirector sqlnulls.NullString `json:"namedirector,omitempty"`
}
