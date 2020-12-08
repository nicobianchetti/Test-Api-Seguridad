package interfaces

import (
	"net/http"

	"github.com/nicobianchetti/Test-api2/src/commons/model"
)

//IPermisoRepository interect with permiso entity
type IPermisoRepository interface {
	Migrate() error
	Create(*model.Permiso) error
	Update(string, *model.Permiso) error
	GetAll() (*model.Permisos, error)
	GetByID(string) (*model.Permiso, error)
	Delete(string) error
}

//IPermisoService interact with IPermisoRepository
type IPermisoService interface {
	Migrate() error
	Create(*model.Permiso) error
	Update(string, *model.Permiso) error
	GetAll() (*model.Permisos, error)
	GetByID(string) (*model.Permiso, error)
	Delete(string) error
}

//IPermisoController interac with IPermisoService
type IPermisoController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
