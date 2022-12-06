package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	requestVacationStorage    storage.IRequestVacationStorage       = storage.NewRequestVacation()
	requestVacationService    service.IRequestVacationService       = service.NewRequestVacationService(requestVacationStorage)
	requestVacationController controller.IRequestVacationController = controller.NewRequestVacationController(requestVacationService)
)

func SetRequestVacationRoutes(router *mux.Router) *mux.Router {
	router = router.PathPrefix("/requestvacations").Subrouter()
	router.HandleFunc("", requestVacationController.Create).Methods("POST")
	router.HandleFunc("", requestVacationController.GetRequestsVacations).Methods("GET")
	router.HandleFunc("/{uuid}", requestVacationController.GetOneRequestVacation).Methods("GET")
	router.HandleFunc("/{uuid}", requestVacationController.UpdateRequestVacation).Methods("PUT")
	router.HandleFunc("/{uuid}", requestVacationController.Delete).Methods("DELETE")
	return router
}
