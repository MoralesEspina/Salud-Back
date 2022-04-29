package models

type Credentials struct {
	ActualPassword string `json:"actual_password,omitempty"`
	NewPassword    string `json:"new_password,omitempty"`
}
