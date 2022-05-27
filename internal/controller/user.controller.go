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

type userController struct{}

var userService service.UserService

// NewUserController retorna un nuevo controller de tipo usuario controller
func NewUserController(service service.UserService) UserController {
	userService = service
	return &userController{}
}

// UserController contiene todos los controladores de usuario
type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetOneUser(w http.ResponseWriter, r *http.Request)
	ManyUsers(w http.ResponseWriter, r *http.Request)
	Rols(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UserInformationByToken(w http.ResponseWriter, r *http.Request)
}

func (*userController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	idinsert, err := userService.Create(r.Context(), &user)

	if err == lib.ErrDuplicateUser {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrDuplicateUser.Error(),
		}, http.StatusConflict)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:       true,
			Message:  "Usuario creado satisfactoriamente",
			IDInsert: idinsert,
		}, http.StatusCreated)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*userController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	idUpdated, err := userService.Update(r.Context(), mux.Vars(r)["uuid"], user)

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

func (*userController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp, err := userService.Login(r.Context(), &user)

	if err == lib.ErrUserNotFound {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrUserNotFound.Error(),
		}, http.StatusNotFound)
		return
	}

	if err == nil {
		respond(w, resp, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*userController) UserInformationByToken(w http.ResponseWriter, r *http.Request) {
	tokenInfo, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrUnauthenticated.Error(),
		}, http.StatusUnauthorized)
		return
	}

	resp, err := userService.UserInformationByToken(r.Context(), tokenInfo.ID)

	if err == lib.ErrUserNotFound {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrUserNotFound.Error(),
		}, http.StatusNotFound)
		return
	}

	if err == nil {
		respond(w, resp, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// func (*userController) Update(w http.ResponseWriter, r *http.Request) {
// 	_, ok := middleware.IsAuthenticated(r.Context())
// 	if !ok {
// 		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
// 		return
// 	}

// 	defer r.Body.Close()
// 	var user models.User

// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		respond(w, response{
// 			Ok:      false,
// 			Message: err.Error(),
// 		}, http.StatusBadRequest)
// 		return
// 	}

// 	vars := mux.Vars(r)

// 	err := userService.Update(r.Context(), vars["id"], user.Rol)
// 	if err == nil {
// 		respond(w, response{
// 			Ok:      true,
// 			Message: "Usuario actualizado satisfactoriamente",
// 		}, http.StatusAccepted)
// 		return
// 	}

// 	if err != nil {
// 		respondError(w, err)
// 		return
// 	}
// }

func (*userController) GetOneUser(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := userService.GetOneUser(r.Context(), vars["uuid"])
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

func (*userController) ManyUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	data, err := userService.ManyUsers(r.Context())
	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:   false,
			Data: users,
		}, http.StatusNotFound)
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

func (*userController) Rols(w http.ResponseWriter, r *http.Request) {

	data, err := userService.Roles(r.Context())
	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:   false,
			Data: data,
		}, http.StatusNotFound)
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

func (*userController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	tokenInfo, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrUnauthenticated.Error(),
		}, http.StatusUnauthorized)
		return
	}
	credentials := models.Credentials{}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err := userService.ChangePassword(r.Context(), tokenInfo.ID, credentials.ActualPassword, credentials.NewPassword)
	if err == lib.ErrUserNotFound {
		respond(w, response{
			Ok:      false,
			Message: "Registro no encontrado",
		}, http.StatusNotFound)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Password modificada correctamente",
		}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	_, err := userService.DeleteUser(r.Context(), uuid)
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
