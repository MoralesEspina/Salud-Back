package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/gorilla/mux"
)

type personController struct{}

var IPersonService service.PersonService

// NewPersonController retorna un nuevo controller de tipo usuario controller
func NewPersonController(personService service.PersonService) PersonController {
	IPersonService = personService
	return &personController{}
}

// PersonController contiene todos los controladores de usuario
type PersonController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	GetMany(w http.ResponseWriter, r *http.Request)
	GetManyWithFullInformation(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)

	CreateSubstitute(w http.ResponseWriter, r *http.Request)
	GetOneSubstitute(w http.ResponseWriter, r *http.Request)
	GetSubstitutes(w http.ResponseWriter, r *http.Request)
	GetNamePerson(w http.ResponseWriter, r *http.Request)
}

func (*personController) Create(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var request models.Person

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := IPersonService.Create(r.Context(), request)
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

func (*personController) GetOne(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := IPersonService.GetOne(r.Context(), vars["uuid"])
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

func (*personController) GetMany(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	limit, err := strconv.Atoi(lib.ValuesURL(r, "limit"))
	if err != nil {
		respond(w, response{
			Ok:      false,
			Message: "Los argumentos enviados por url son invalidos",
		}, http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(lib.ValuesURL(r, "page"))
	filter := lib.ValuesURL(r, "filter")
	if err != nil {
		respond(w, response{
			Ok:      false,
			Message: "Los argumentos enviados por url son invalidos",
		}, http.StatusBadRequest)
		return
	}

	data, err := IPersonService.GetMany(r.Context(), filter, page, limit)
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

func (*personController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person models.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	idUpdated, err := IPersonService.Update(r.Context(), mux.Vars(r)["uuid"], person)

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

func (*personController) GetManyWithFullInformation(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	limit, err := strconv.Atoi(lib.ValuesURL(r, "limit"))
	if err != nil {
		respond(w, response{
			Ok:      false,
			Message: "Los argumentos enviados por url son invalidos",
		}, http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(lib.ValuesURL(r, "page"))
	filter := lib.ValuesURL(r, "filter")
	if err != nil {
		respond(w, response{
			Ok:      false,
			Message: "Los argumentos enviados por url son invalidos",
		}, http.StatusBadRequest)
		return
	}

	data, err := IPersonService.GetManyWithFullInformation(r.Context(), filter, offset, limit)

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

func ValidateExistePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	data, err := service.ValidateExistePerson(r.Context(), vars["uuid"])
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

func (*personController) GetNamePerson(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := IPersonService.GetSubstitutes(r.Context())
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
