package controller

import (
	"encoding/json"
	"net/http"

	//"strconv"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/gorilla/mux"
)

type personEducationController struct{}

var IPersonEducationService service.PersonEducationService

// NewPersonEducationController retorna un nuevo controller de tipo usuario controller
func NewPersonEducationController(personEducationService service.PersonEducationService) PersonEducationController {
	IPersonEducationService = personEducationService
	return &personEducationController{}
}

// PersonEducationController contiene todos los controladores de usuario
type PersonEducationController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetEducations(w http.ResponseWriter, r *http.Request)
}

func (*personEducationController) Create(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}
	defer r.Body.Close()
	var request models.PersonEducation

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := IPersonEducationService.Create(r.Context(), request)
	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Registro creado satisfactoriamente",
			Data:    result.UUID,
		}, http.StatusCreated)
		return
	}

	if err != nil {
		if lib.DecodeMySQLError(err).Number == 1452 {
			respond(w, response{
				Ok:      false,
				Message: lib.Err1452,
				Data:    result,
			}, http.StatusBadRequest)
			return
		}

		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*personEducationController) GetEducations(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := IPersonEducationService.GetEducations(r.Context(), vars["uuid"])
	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:      false,
			Data:    data,
			Message: lib.ErrNotFound.Error(),
		}, http.StatusNotFound)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:   true,
			Data: data,
		}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
