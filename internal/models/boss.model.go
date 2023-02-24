package models

import "github.com/Mynor2397/sqlnulls"

// Person implenta el modelo de la base de datos
type Boss struct {
	ID   sqlnulls.NullString `json:"id,omitempty"`
	Name sqlnulls.NullString `json:"name,omitempty"`
}
