package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	userStorage    storage.UserStorage       = storage.NewUserStorage()
	userService    service.UserService       = service.NewUserService(userStorage)
	userController controller.UserController = controller.NewUserController(userService)
)

// SetUserRoutes registra la rutas a usar para los controladires de usuario
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", userController.Login).Methods("POST")
	router.HandleFunc("/user/verify", userController.UserInformationByToken).Methods("GET")

	user := router.PathPrefix("/users").Subrouter()
	user.Use(middleware.AuthForAmdmin)
	user.HandleFunc("/register", userController.Create).Methods("POST")
	user.HandleFunc("", userController.ManyUsers).Methods("GET")
	user.HandleFunc("/rols", userController.Rols).Methods("GET")
	user.HandleFunc("/changepassword", userController.ChangePassword).Methods("POST")

	return router
}
