package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	jobStorage    storage.IJobStorage       = storage.NewJobStorage()
	jobService    service.IJobService       = service.NewJobService(jobStorage)
	jobController controller.IJobController = controller.NewJobController(jobService)
)

// SetJobRoutes registra la rutas a usar para los controladires de usuario
func SetJobRoutes(router *mux.Router) *mux.Router {

	job := router.PathPrefix("/jobs").Subrouter()
	job.Use(middleware.AuthForAmdmin)
	job.HandleFunc("", jobController.ManyJobs).Methods("GET")
	job.HandleFunc("", jobController.CreateJob).Methods("POST")
	job.HandleFunc("/{uuid}", jobController.OneJob).Methods("GET")
	job.HandleFunc("/{uuid}", jobController.Delete).Methods("DELETE")
	job.HandleFunc("/{uuid}", jobController.Update).Methods("PUT")

	return router
}
