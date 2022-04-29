package cross

import (
	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/mysql"
)

// GenerateDynamicNumberRegister genera el numero de registro basado en el campo submittedAt
// el cual contiene fecha y hora de registro
func GenerateDynamicNumberRegister(table string) (int, error) {
	db := mysql.Connect()
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM "+table+" WHERE YEAR(submittedAt) = YEAR(?)", lib.TimeZone("America/Guatemala").DateTime).Scan(&count)

	if err != nil {
		return 0, err
	}

	if count == 0 {
		count = 1
	} else {
		count += 1
	}

	return count, nil
}
