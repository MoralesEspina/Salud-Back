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

type workDependencyController struct{}

var WorkDependencyService service.IWorkDependencyService

// NewWorkDependencyController retorna un nuevo controller de tipo usuario controller
func NewWorkDependencyController(workdependencyService service.IWorkDependencyService) IWorkDependencyController {
	WorkDependencyService = workdependencyService
	return &workDependencyController{}
}

// IWorkDependencyController contiene todos los controladores de usuario
type IWorkDependencyController interface {
	GetManyWorks(w http.ResponseWriter, r *http.Request)
	CreateWorkDependency(w http.ResponseWriter, r *http.Request)
	OneWorkDependency(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func (*workDependencyController) GetManyWorks(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := WorkDependencyService.GetManyWorkDependency(r.Context())
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

func (*workDependencyController) CreateWorkDependency(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var workDependency models.WorkDependency

	if err := json.NewDecoder(r.Body).Decode(&workDependency); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := WorkDependencyService.CreateWorkDependency(r.Context(), workDependency)
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

func (*workDependencyController) OneWorkDependency(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := WorkDependencyService.OneWorkDependency(r.Context(), mux.Vars(r)["uuid"])
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

func (*workDependencyController) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	err := WorkDependencyService.Delete(r.Context(), uuid)
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

func (*workDependencyController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var workDependency models.WorkDependency

	if err := json.NewDecoder(r.Body).Decode(&workDependency); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
			Data:    emptyArray,
		}, http.StatusBadRequest)
		return
	}

	result, err := WorkDependencyService.Update(r.Context(), workDependency, mux.Vars(r)["uuid"])
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
