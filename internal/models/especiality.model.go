package models

import "github.com/Mynor2397/sqlnulls"

type Especiality struct {
	UUIDEspeciality sqlnulls.NullString `json:"uuid_especiality,omitempty"`
	Name            sqlnulls.NullString `json:"name,omitempty"`
	Description     string              `json:"description,omitempty"`
}
