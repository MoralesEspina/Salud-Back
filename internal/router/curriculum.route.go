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

	referencesStorage    storage.ReferencesStorage       = storage.NewReferencesStorage()
	referencesService    service.ReferencesService       = service.NewReferencesService(referencesStorage)
	referencesController controller.ReferencesController = controller.NewReferencesController(referencesService)

	personEducationStorage    storage.PersonEducationStorage       = storage.NewPersonEducationStorage()
	personEducationService    service.PersonEducationService       = service.NewPersonEducationService(personEducationStorage)
	personEducationController controller.PersonEducationController = controller.NewPersonEducationController(personEducationService)

	workExpStorage    storage.WorkExpStorage       = storage.NewWorkExpStorage()
	workExpService    service.WorkExpService       = service.NewWorkExpService(workExpStorage)
	workExpController controller.WorkExpController = controller.NewWorkExpController(workExpService)
)

// SetCurriculumRoutes registra la rutas a usar para los controladires de usuario
func SetCurriculumRoutes(router *mux.Router) {
	router.Use(middleware.Loger)
	curriculum := router.PathPrefix("/curriculums").Subrouter()
	// person.Use(middleware.AuthForAmdmin)
	curriculum.Handle("", middleware.AuthForAmdminTypeHTTP(curriculumController.Create)).Methods("POST")
	curriculum.HandleFunc("/{uuid}", curriculumController.GetOne).Methods("GET")
	curriculum.Handle("/{uuid}", middleware.AuthForAmdminTypeHTTP(curriculumController.Update)).Methods("PUT")

	references := router.PathPrefix("/references").Subrouter()
	// person.Use(middleware.AuthForAmdmin)
	references.Handle("", middleware.AuthForAmdminTypeHTTP(referencesController.Create)).Methods("POST")
	references.HandleFunc("/{uuid}", referencesController.GetReferences).Methods("GET")
	references.HandleFunc("/{uuid}", referencesController.DeleteReferences).Methods("DELETE")

	personEducation := router.PathPrefix("/personEducation").Subrouter()
	// person.Use(middleware.AuthForAmdmin)
	personEducation.Handle("", middleware.AuthForAmdminTypeHTTP(personEducationController.Create)).Methods("POST")
	personEducation.HandleFunc("/{uuid}", personEducationController.GetEducations).Methods("GET")
	personEducation.HandleFunc("/{uuid}", personEducationController.DeleteEducations).Methods("DELETE")

	workExp := router.PathPrefix("/workExp").Subrouter()
	// person.Use(middleware.AuthForAmdmin)
	workExp.Handle("", middleware.AuthForAmdminTypeHTTP(workExpController.Create)).Methods("POST")
	workExp.HandleFunc("/{uuid}", workExpController.GetWorks).Methods("GET")
	workExp.HandleFunc("/{uuid}", workExpController.DeleteWorks).Methods("DELETE")
}
