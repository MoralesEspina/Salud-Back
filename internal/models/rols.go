package models

type Rol struct {
	IDRol string `json:"id_rol"`
	TypeRol string `json:"type_rol"`
	Description string `json:"description,omitempty"`
}