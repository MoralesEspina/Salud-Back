package router

import (
	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/gorilla/mux"
)

var (
	personStorage    storage.PersonStorage       = storage.NewPersonStorage()
	personService    service.PersonService       = service.NewPersonService(personStorage)
	personController controller.PersonController = controller.NewPersonController(personService)
)

// SetPersonRoutes registra la rutas a usar para los controladires de usuario
func SetPersonRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/validation/certify/{uuid}", controller.ValidateExistePerson).Methods("GET")
	person := router.PathPrefix("/persons").Subrouter()
	// person.Use(middleware.AuthForAmdmin)
	person.Handle("", middleware.AuthForAmdminTypeHTTP(personController.Create)).Methods("POST")
	person.HandleFunc("/{uuid}", personController.GetOne).Methods("GET")
	person.HandleFunc("", personController.GetMany).Methods("GET")
	person.Handle("/{uuid}", middleware.AuthForAmdminTypeHTTP(personController.Update)).Methods("PUT")

	router.HandleFunc("/substitutes", personController.CreateSubstitute).Methods("POST")
	router.HandleFunc("/substitutes", personController.GetSubstitutes).Methods("GET")
	router.HandleFunc("/substitutes/{uuid}", personController.GetOneSubstitute).Methods("GET")

	person.Handle("/information/full", middleware.AuthForAmdminTypeHTTP(personController.GetManyWithFullInformation)).Methods("GET")
	return router
}
