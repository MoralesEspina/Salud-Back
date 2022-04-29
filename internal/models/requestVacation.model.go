package models

type RequestVacation struct {
	UUIDRequestVacation string `json:"uuid_request_vacation,omitempty"`
	Register            int    `json:"register,omitempty"`
	SubmittedAt         string `json:"submitted_at,omitempty"`
	ModifiedAt          string `json:"modified_at,omitempty"`
	LastYearVacation    string `json:"last_year_vacation,omitempty"`
	VacationYearRequest string `json:"vacation_year_request,omitempty"`
	LastVacationFrom    string `json:"last_vacation_from,omitempty"`
	LastVacationTo      string `json:"last_vacation_to,omitempty"`
	VacationFromDate    string `json:"vacation_from_date,omitempty"`
	VacationToDate      string `json:"vacation_to_date,omitempty"`
	HasVacationDay      bool   `json:"has_vacation_day"`
	DaysQuantity        int    `json:"days_quantity"`
	Observations        string `json:"observations,omitempty"`
	Person              `json:"person,omitempty"`
	PersonServer        Person `json:"person_server,omitempty"`
	UUIDUser            string `json:"uuid_user,omitempty"`
}
