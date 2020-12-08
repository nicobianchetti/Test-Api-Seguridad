package permiso

import (
	"time"

	"github.com/google/uuid"
	"github.com/nicobianchetti/Test-api2/src/commons/interfaces"
	"github.com/nicobianchetti/Test-api2/src/commons/model"
)

type permisoService struct {
	service interfaces.IPermisoRepository
}

//NewPermisoService create new instance of service
func NewPermisoService(rPermiso interfaces.IPermisoRepository) interfaces.IPermisoService {
	return &permisoService{rPermiso}
}

// Migrate is used for migrate permiso
func (s *permisoService) Migrate() error {
	return s.service.Migrate()
}

// Create is used for create a permiso
func (s *permisoService) Create(p *model.Permiso) error {
	p.ID = uuid.New().String()
	p.Status = true
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return s.service.Create(p)
}

// GetAll is used for get all the permisos
func (s *permisoService) GetAll() (*model.Permisos, error) {
	return s.service.GetAll()
}

// GetByID is used for get a permiso
func (s *permisoService) GetByID(id string) (*model.Permiso, error) {
	return s.service.GetByID(id)
}

// Update is used for update a permiso
func (s *permisoService) Update(id string, p *model.Permiso) error {

	p.UpdatedAt = time.Now()

	return s.service.Update(id, p)
}

// Delete is used for delete a permiso
func (s *permisoService) Delete(id string) error {
	return s.service.Delete(id)
}
