package models

// Config es la configuracion del servidor
type Config struct {
	PORT          string
	HOSTDB        string
	PORTDB        int
	USERDB        string
	PASSWORDDB    string
	DATABASE      string
	DEVHOSTDB     string
	DEVPORTDB     int
	DEVUSERDB     string
	DEVPASSWORDDB string
	DEVDATABASE   string
	PRODUCTION    bool
	USESSL        bool
}
