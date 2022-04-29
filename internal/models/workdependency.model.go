package models

import "github.com/Mynor2397/sqlnulls"

type WorkDependency struct {
	UUIDWork sqlnulls.NullString `json:"uuid_work,omitempty"`
	Name     sqlnulls.NullString `json:"name,omitempty"`
}
