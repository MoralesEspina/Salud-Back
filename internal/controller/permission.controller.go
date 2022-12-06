package controller

import (
	"encoding/json"
	"net/http"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/gorilla/mux"
)

type permissionController struct{}

var IPermissionService service.IPermissionService

func NewPermissionController(permissionService service.IPermissionService) IPermissionController {
	IPermissionService = permissionService
	return &permissionController{}
}

type IPermissionController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetPermissions(w http.ResponseWriter, r *http.Request)
	GetOnePermission(w http.ResponseWriter, r *http.Request)
	UpdatePermission(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (*permissionController) Create(w http.ResponseWriter, r *http.Request) {
	tokenInfo, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var request models.Permission

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := IPermissionService.Create(r.Context(), request, tokenInfo.ID)
	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Registro creado satisfactoriamente",
			Data:    result.Uuid,
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

func (*permissionController) GetPermissions(w http.ResponseWriter, r *http.Request) {
	tokenInfo, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := IPermissionService.GetPermissions(r.Context(), tokenInfo.ID, tokenInfo.Rol)
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

func (*permissionController) GetOnePermission(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := IPermissionService.GetOnePermission(r.Context(), vars["uuid"])
	if err == lib.ErrSQL404 {
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

func (*permissionController) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var request models.Permission

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	result, err := IPermissionService.UpdatePermission(r.Context(), request, mux.Vars(r)["uuid"])
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

func (*permissionController) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	_, err := IPermissionService.DeletePermission(r.Context(), uuid)
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