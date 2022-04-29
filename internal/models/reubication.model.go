package models

import "github.com/Mynor2397/sqlnulls"

type Reubication struct {
	UUIDReubication sqlnulls.NullString `json:"uuid_reubication,omitempty"`
	Name            sqlnulls.NullString `json:"name,omitempty"`
	Description     string              `json:"description,omitempty"`
}
