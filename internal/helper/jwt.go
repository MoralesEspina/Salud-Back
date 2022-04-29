package helper

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

var (
	// PrivateKey es la llave primaria para firmar el token
	PrivateKey *rsa.PrivateKey

	// PublicKey es la llave publica para firmar el token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./internal/certificates/private.rsa")
	if err != nil {
		log.Fatal("Error al leer el archivo: ", err)
	}

	publicBytes, err := ioutil.ReadFile("./internal/certificates/public.rsa.pub")
	if err != nil {
		log.Fatal("Error al leer el archivo: ", err)
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Error al intentar parsear la llave privada")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Error al intentar parsear la llave publica")
	}

}

// GenerateJWT genera un token JWT
func GenerateJWT(user *models.User) string {
	claims := models.Claim{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
			Issuer:    "ministerio de salud",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Println("No se ha podido firmar el token: ", err)
	}

	return result
}

// AuthUserID evalua el token y devuelve un id y un error
func AuthUserID(r *http.Request) (string, string, error) {
	claims := &models.Claim{}
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, claims, func(token *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			typeError := err.(*jwt.ValidationError)

			switch typeError.Errors {
			case jwt.ValidationErrorExpired:
				return "", "", lib.ErrTokenExpired

			case jwt.ValidationErrorSignatureInvalid:
				return "", "", lib.ErrInvalidsignature

			default:
				return "", "", err
			}
		}
	}
	var (
		id  string
		rol string
	)

	if token.Valid {
		id = claims.ID
		rol = claims.Rol
	}
	return id, rol, nil
}
