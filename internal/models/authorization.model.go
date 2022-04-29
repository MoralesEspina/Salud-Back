package models

import "github.com/Mynor2397/sqlnulls"

// Authorization implementa el modelo de autorización de la base de datos
type Authorization struct {
	UUIDAuthorization  string `json:"uuid_authorization,omitempty"`
	Register           int    `json:"register,omitempty"`
	SubmittedAt        string `json:"submitted_at,omitempty"`
	ModifiedAt         string `json:"modified_at,omitempty"`
	Startdate          string `json:"startdate,omitempty"`
	Enddate            string `json:"enddate,omitempty"`
	Resumework         string `json:"resumework,omitempty"`
	Holidays           int    `json:"holidays,omitempty"`
	TotalDays          int    `json:"total_days,omitempty"`
	Pendingdays        int    `json:"pendingdays,omitempty"`
	Observation        string `json:"observation,omitempty"`
	Authorizationyear  string `json:"authorizationyear,omitempty"`
	Workdependency     string `json:"workdependency,omitempty"`
	WorkdependencyUUID string `json:"workdependency_uuid,omitempty"`
	User               string `json:"user,omitempty"`

	Person `json:"person,omitempty"`

	PersonnelOfficer          sqlnulls.NullString `json:"personnelOfficer,omitempty"`
	PersonnelOfficerPosition  sqlnulls.NullString `json:"personnelOfficerPosition,omitempty"`
	PersonnelOfficerArea      sqlnulls.NullString `json:"personnelOfficerArea,omitempty"`
	ExecutiveDirector         sqlnulls.NullString `json:"executiveDirector,omitempty"`
	ExecutiveDirectorPosition sqlnulls.NullString `json:"executiveDirectorPosition,omitempty"`
	ExecutiveDirectorArea     sqlnulls.NullString `json:"executiveDirectorArea,omitempty"`
}
