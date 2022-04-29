package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	workDependencyStorage    storage.IWorkDependencyStorage       = storage.NewWorkDependencyStorage()
	workDependencyService    service.IWorkDependencyService       = service.NewWorkDependencyService(workDependencyStorage)
	workDependencyController controller.IWorkDependencyController = controller.NewWorkDependencyController(workDependencyService)
)

// SetWorkDependencyRoutes registra la rutas a usar para los controladires de usuario
func SetWorkDependencyRoutes(router *mux.Router) *mux.Router {

	workDependency := router.PathPrefix("/works").Subrouter()
	workDependency.Use(middleware.AuthForAmdmin)
	workDependency.HandleFunc("", workDependencyController.GetManyWorks).Methods("GET")
	workDependency.HandleFunc("", workDependencyController.CreateWorkDependency).Methods("POST")
	workDependency.HandleFunc("/{uuid}", workDependencyController.OneWorkDependency).Methods("GET")
	workDependency.HandleFunc("/{uuid}", workDependencyController.Delete).Methods("DELETE")
	workDependency.HandleFunc("/{uuid}", workDependencyController.Update).Methods("PUT")

	return router
}
