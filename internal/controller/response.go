package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var emptyArray = [...]string{}

// var emptyObject = make(map[string]string)

type response struct {
	Ok       bool        `json:"ok"`
	Message  string      `json:"message,omitempty"`
	IDInsert string      `json:"id_insert,omitempty"`
	Data     interface{} `json:"data"`
}

func respond(w http.ResponseWriter, v interface{}, statuscode int) {
	b, err := json.Marshal(v)

	if err != nil {
		respondError(w, fmt.Errorf("No se puede obtener la respuesta: %v", err))
		return
	}

	w.WriteHeader(statuscode)
	w.Write(b)

}

func respondError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
