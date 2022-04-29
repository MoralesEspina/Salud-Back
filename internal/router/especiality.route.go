package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	especialityStorage    storage.IEspecialityStorage       = storage.NewEspecialityStorage()
	especialityService    service.IEspecialityService       = service.NewEspecialityService(especialityStorage)
	especialityController controller.IEspecialityController = controller.NewEspecialityController(especialityService)
)

// SetEspecialityRoutes registra la rutas a usar para los controladires de usuario
func SetEspecialityRoutes(router *mux.Router) *mux.Router {

	especiality := router.PathPrefix("/especialities").Subrouter()
	especiality.Use(middleware.AuthForAmdmin)
	especiality.HandleFunc("", especialityController.Especialities).Methods("GET")
	especiality.HandleFunc("", especialityController.CreateEspeciality).Methods("POST")
	especiality.HandleFunc("/{uuid}", especialityController.Especality).Methods("GET")
	especiality.HandleFunc("/{uuid}", especialityController.Delete).Methods("DELETE")
	especiality.HandleFunc("/{uuid}", especialityController.Update).Methods("PUT")

	return router
}
