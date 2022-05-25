package router

import (
	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/gorilla/mux"
)

var (
	curriculumStorage    storage.CurriculumStorage       = storage.NewCurriculumStorage()
	curriculumService    service.CurriculumService       = service.NewCurriculumService(curriculumStorage)
	curriculumController controller.CurriculumController = controller.NewCurriculumController(curriculumService)
)

// SetCurriculumRoutes registra la rutas a usar para los controladires de usuario
func SetCurriculumRoutes(router *mux.Router) *mux.Router {
	curriculum := router.PathPrefix("/curriculums").Subrouter()
	// person.Use(middleware.AuthForAmdmin)
	curriculum.Handle("", middleware.AuthForAmdminTypeHTTP(curriculumController.Create)).Methods("POST")
	curriculum.HandleFunc("/{uuid}", curriculumController.GetOne).Methods("GET")
	return router
}
