package permiso

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicobianchetti/Test-api2/src/commons/interfaces"
)

//permisoRouter struct content functions of controller Permiso
type permisoRouter struct {
	cPermiso interfaces.IPermisoController
}

//NewPermisoRouter create the new instanst of Router Permiso
func NewPermisoRouter(subRouter *mux.Router, cPermiso interfaces.IPermisoController) {
	rPermiso := permisoRouter{cPermiso}
	rPermiso.Routes(subRouter)
}

//NewPermisoRouterTest create the new instanst of Router Permiso test
func NewPermisoRouterTest(subRouter *mux.Router, cPermiso interfaces.IPermisoController) permisoRouter {
	rPermiso := permisoRouter{cPermiso}
	rPermiso.Routes(subRouter)
	return rPermiso
}

//Routes define routes for permiso
func (r *permisoRouter) Routes(subRouter *mux.Router) {
	subRouter.
		Path("/permisos").
		Handler(http.HandlerFunc(r.cPermiso.GetAll)).
		Methods(http.MethodGet)
	subRouter.
		Path("/permiso/{id}").
		Handler(http.HandlerFunc(r.cPermiso.GetByID)).
		Methods(http.MethodGet)
	subRouter.
		Path("/permiso").
		Handler(http.HandlerFunc(r.cPermiso.Create)).
		Methods(http.MethodPost)
	subRouter.
		Path("/permiso/{id}").
		Handler(http.HandlerFunc(r.cPermiso.Update)).
		Methods(http.MethodPut)
	subRouter.
		Path("/permiso/{id}").
		Handler(http.HandlerFunc(r.cPermiso.Update)).
		Methods(http.MethodPatch)
	subRouter.
		Path("/permiso/{id}").
		Handler(http.HandlerFunc(r.cPermiso.Delete)).
		Methods(http.MethodDelete)
}
