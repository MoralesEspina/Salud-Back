package models

type Rols struct {
	ID          string `json:"id"`
	Role        string `json:"role,omitempty"`
	Description string `json:"description,omitempty"`
}
