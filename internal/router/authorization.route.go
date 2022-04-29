package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	authorizationStorage    storage.AuthorizationStorage       = storage.NewAuthorizationStorage()
	authorizationService    service.AuthorizationService       = service.NewAuthorizationService(authorizationStorage)
	authorizationController controller.AuthorizationController = controller.NewAuthorizationController(authorizationService)
)

// SetAuthorizationRoutes registra la rutas a usar para los controladires de usuario
func SetAuthorizationRoutes(router *mux.Router) *mux.Router {

	authorization := router.PathPrefix("/authorizations").Subrouter()
	authorization.Use(middleware.AuthForAmdmin)
	authorization.HandleFunc("", authorizationController.Create).Methods("POST")
	authorization.HandleFunc("", authorizationController.GetManyAuthorizations).Methods("GET")
	authorization.HandleFunc("/{uuid}", authorizationController.GetOnlyAuthorization).Methods("GET")
	authorization.HandleFunc("/{uuid}", authorizationController.UpdateAuthorization).Methods("PUT")
	authorization.HandleFunc("/pdfauthorization/{uuid}", authorizationController.GetOnlyAuthorizationPDF).Methods("GET")

	router.HandleFunc("/reports/authorizations", authorizationController.VacationsReport)

	return router
}
