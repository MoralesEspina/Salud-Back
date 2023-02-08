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

type curriculumController struct{}

var ICurriculumService service.CurriculumService

// NewCurriculumController retorna un nuevo controller de tipo usuario controller
func NewCurriculumController(curriculumService service.CurriculumService) CurriculumController {
	ICurriculumService = curriculumService
	return &curriculumController{}
}

// CurriculumController contiene todos los controladores de usuario
type CurriculumController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (*curriculumController) Create(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var request models.Curriculum

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := ICurriculumService.Create(r.Context(), request)
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

func (*curriculumController) GetOne(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := ICurriculumService.GetOne(r.Context(), vars["uuid"])
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

func (*curriculumController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var curriculum models.Curriculum

	if err := json.NewDecoder(r.Body).Decode(&curriculum); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	idUpdated, err := ICurriculumService.Update(r.Context(), mux.Vars(r)["uuid"], curriculum)

	if err == nil {
		respond(w, response{
			Ok:       true,
			Message:  "Registro actualizado satisfactoriamente",
			IDInsert: idUpdated,
		}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*curriculumController) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	_, err := ICurriculumService.DeleteCurriculum(r.Context(), uuid)
	if err != nil {
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
