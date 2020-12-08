package permiso

import (
	"os"
	"testing"

	"github.com/gorilla/mux"
)

const (
	port = "8080"
	ip   = ""
)

var roPermiso permisoRouter
var r *mux.Router

func TestMain(m *testing.M) {
	// r = mux.NewRouter()

	// driver := database.Postgres
	// database.New(driver)
	// db := database.DB()

	// rPermiso := NewPermisoRepository(db)
	// sPermiso := NewPermisoService(rPermiso)
	// cPermiso := NewPermisoController(sPermiso)
	// roPermiso = NewPermisoRouterTest(r.PathPrefix("/v1/permiso").Subrouter(), cPermiso)

	code := m.Run()
	os.Exit(code)
}
