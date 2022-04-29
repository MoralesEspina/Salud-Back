package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Claim es el modelo de reclamaciones
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
