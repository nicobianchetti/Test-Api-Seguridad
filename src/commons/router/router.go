package router

import (
	"github.com/gorilla/mux"
	"github.com/nicobianchetti/Test-api2/database"
	"github.com/nicobianchetti/Test-api2/src/commons/permiso"
)

/*
SetupRoutesPermiso calls all server endpoints
*/
func SetupRoutesPermiso(muxRouter *mux.Router) {
	/*
		Defino driver de base de datos
	*/
	driver := database.Postgres
	database.New(driver)
	db := database.DB()

	/*
		Conexión entidad Permiso
	*/
	rPermiso := permiso.NewPermisoRepository(db)
	sPermiso := permiso.NewPermisoService(rPermiso)
	cPermiso := permiso.NewPermisoController(sPermiso)
	permiso.NewPermisoRouter(muxRouter.PathPrefix("/v1/permiso").Subrouter(), cPermiso)
}

/*
SetupRoutesPermisoTest calls all server endpoints
*/
// func SetupRoutesPermisoTest(muxRouter *mux.Router, db *gorm.DB) interfaces.IPermisoController {
// 	rPermiso := permiso.NewPermisoRepository(db)
// 	sPermiso := permiso.NewPermisoService(rPermiso)
// 	cPermiso := permiso.NewPermisoController(sPermiso)
// 	permiso.NewPermisoRouter(muxRouter.PathPrefix("/v1/permiso").Subrouter(), cPermiso)

// 	return cPermiso
// }
