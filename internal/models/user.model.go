package models

// User es el modelo de características de un usuario
type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Rol      string `json:"rol,omitempty"`
	IDRol    int64  `json:"id_rol,omitempty"`
	Token    string `json:"token,omitempty"`
}
