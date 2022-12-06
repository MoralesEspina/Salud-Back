package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	permissionStorage    storage.IPermissionStorage       = storage.NewPermission()
	permissionService    service.IPermissionService       = service.NewPermissionService(permissionStorage)
	permissionController controller.IPermissionController = controller.NewPermissionController(permissionService)
)

func SetPermissionRoutes(router *mux.Router) *mux.Router {
	router = router.PathPrefix("/permission").Subrouter()
	router.HandleFunc("", permissionController.Create).Methods("POST")
	router.HandleFunc("", permissionController.GetPermissions).Methods("GET")
	router.HandleFunc("/{uuid}", permissionController.GetOnePermission).Methods("GET")
	router.HandleFunc("/{uuid}", permissionController.UpdatePermission).Methods("PUT")
	router.HandleFunc("/{uuid}", permissionController.Delete).Methods("DELETE")
	return router
}
