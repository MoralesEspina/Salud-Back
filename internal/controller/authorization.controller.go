package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/service"
)

type authorizationController struct{}

var AuthorizationService service.AuthorizationService

// NewAuthorizationController retorna un nuevo controller de tipo usuario controller
func NewAuthorizationController(authorizationService service.AuthorizationService) AuthorizationController {
	AuthorizationService = authorizationService
	return &authorizationController{}
}

// AuthorizationController contiene todos los controladores de usuario
type AuthorizationController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetManyAuthorizations(w http.ResponseWriter, r *http.Request)
	GetOnlyAuthorization(w http.ResponseWriter, r *http.Request)
	UpdateAuthorization(w http.ResponseWriter, r *http.Request)

	GetOnlyAuthorizationPDF(w http.ResponseWriter, r *http.Request)
	VacationsReport(w http.ResponseWriter, r *http.Request)
}

func (*authorizationController) Create(w http.ResponseWriter, r *http.Request) {
	infoToken, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var authorization models.Authorization

	if err := json.NewDecoder(r.Body).Decode(&authorization); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	authorization.User = infoToken.ID

	result, err := AuthorizationService.Create(r.Context(), authorization)
	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Registro creado satisfactoriamente",
			Data:    result,
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

func (*authorizationController) GetManyAuthorizations(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := AuthorizationService.GetManyAuthorizations(r.Context())
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

func (*authorizationController) GetOnlyAuthorization(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := AuthorizationService.GetOnlyAuthorization(r.Context(), vars["uuid"])
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

func (*authorizationController) UpdateAuthorization(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var authorization models.Authorization

	if err := json.NewDecoder(r.Body).Decode(&authorization); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := AuthorizationService.UpdateAuthorization(r.Context(), authorization, mux.Vars(r)["uuid"])
	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Registro acttualizado satisfactoriamente",
			Data:    result,
		}, http.StatusCreated)
		return
	}

	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:      true,
			Message: "Registro no encontrado",
		}, http.StatusNotFound)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*authorizationController) GetOnlyAuthorizationPDF(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := AuthorizationService.GetOnlyAuthorizationPDF(r.Context(), vars["uuid"])
	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:      false,
			Data:    data,
			Message: lib.ErrNotFound.Error(),
		}, http.StatusOK)
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

func (*authorizationController) VacationsReport(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	startDate := lib.ValuesURL(r, "startdate")
	endDate := lib.ValuesURL(r, "enddate")
	data, err := AuthorizationService.VacationsReport(r.Context(), startDate, endDate)

	if err != nil {

		if err == lib.ErrNotFound {
			respond(w, response{
				Ok:      false,
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
			Data: data,
		}, http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
