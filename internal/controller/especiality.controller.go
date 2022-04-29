package controller

import (
	"encoding/json"
	"net/http"

	// "github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/gorilla/mux"
)

type especialityController struct{}

var EspecialityService service.IEspecialityService

// NewAuthorizationController retorna un nuevo controller de tipo usuario controller
func NewEspecialityController(especialityService service.IEspecialityService) IEspecialityController {
	EspecialityService = especialityService
	return &especialityController{}
}

// IAuthorizationController contiene todos los controladores de usuario
type IEspecialityController interface {
	Especialities(w http.ResponseWriter, r *http.Request)
	Especality(w http.ResponseWriter, r *http.Request)
	CreateEspeciality(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func (*especialityController) CreateEspeciality(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var especiality models.Especiality

	if err := json.NewDecoder(r.Body).Decode(&especiality); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := EspecialityService.CreateEspeciality(r.Context(), especiality)
	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Registro creado satisfactoriamente",
			Data:    result,
		}, http.StatusCreated)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*especialityController) Especialities(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := EspecialityService.Especialities(r.Context())
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

func (*especialityController) Especality(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := EspecialityService.OneEspeciality(r.Context(), mux.Vars(r)["uuid"])
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

func (*especialityController) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	err := EspecialityService.Delete(r.Context(), uuid)
	if err != nil {

		if err.Error() == lib.Status1451.Error() {
			respond(w, response{
				Ok:      false,
				Data:    emptyArray,
				Message: lib.Err1451,
			}, http.StatusNotAcceptable)
			return
		}

		if err == lib.ErrNotFound {
			respond(w, response{
				Ok:      false,
				Data:    emptyArray,
				Message: lib.ErrNotFound.Error(),
			}, http.StatusNotFound)
			return
		}

		respondError(w, err)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:   true,
			Data: uuid,
		}, http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*especialityController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var especiality models.Especiality

	if err := json.NewDecoder(r.Body).Decode(&especiality); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
			Data:    emptyArray,
		}, http.StatusBadRequest)
		return
	}

	result, err := EspecialityService.Update(r.Context(), especiality, mux.Vars(r)["uuid"])
	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:      false,
			Data:    emptyArray,
			Message: lib.ErrNotFound.Error(),
		}, http.StatusNotFound)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Registro actualizado satisfactoriamente",
			Data:    result,
		}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
