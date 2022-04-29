package models

import "github.com/Mynor2397/sqlnulls"

type Job struct {
	UUIDJob     sqlnulls.NullString `json:"uuid_job,omitempty", validate:"eq=36"`
	Name        sqlnulls.NullString `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
}
