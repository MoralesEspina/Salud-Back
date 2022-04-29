package lib

import (
	"encoding/json"
	"log"
	"os"

	"github.com/DasJalapa/reportes-salud/internal/models"
)

// Config del servidor
func Config() models.Config {
	var parameters models.Config
	configfile, err := os.Open("./config/config.json")
	if err != nil {
		log.Println(err.Error())
	}

	defer configfile.Close()

	var configDecoder *json.Decoder = json.NewDecoder(configfile)

	err = configDecoder.Decode(&parameters)
	if err != nil {
		log.Fatal(err.Error())
	}

	if !parameters.PRODUCTION {
		parameters.HOSTDB = parameters.DEVHOSTDB
		parameters.PORTDB = parameters.DEVPORTDB
		parameters.USERDB = parameters.DEVUSERDB
		parameters.PASSWORDDB = parameters.DEVPASSWORDDB
		parameters.DATABASE = parameters.DEVDATABASE
	}
	return parameters
}
