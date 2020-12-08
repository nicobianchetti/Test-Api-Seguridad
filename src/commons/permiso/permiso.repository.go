package permiso

import (
	"github.com/jinzhu/gorm"
	"github.com/nicobianchetti/Test-api2/src/commons/interfaces"
	"github.com/nicobianchetti/Test-api2/src/commons/model"
)

type permisoRepository struct {
	db *gorm.DB
}

//NewPermisoRepository receives a DB connection and creates a new Repository.
func NewPermisoRepository(gormDB *gorm.DB) interfaces.IPermisoRepository {
	return &permisoRepository{gormDB}
}

func (r *permisoRepository) Migrate() error {
	err := r.db.AutoMigrate(&model.Permiso{}).Error
	if err != nil {
		return err
	}

	// fmt.Println("Migración de permiso ejecutada correctamente")
	return nil
}

func (r *permisoRepository) Create(pr *model.Permiso) error {
	err := r.db.Create(&pr).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *permisoRepository) GetAll() (*model.Permisos, error) {

	permisos := make(model.Permisos, 0)
	err := r.db.Find(&permisos).Error

	if err != nil {
		return nil, err
	}

	return &permisos, nil
}

func (r *permisoRepository) GetByID(ID string) (*model.Permiso, error) {

	permiso := model.Permiso{}

	err := r.db.Where("id = ?", ID).First(&permiso).Error

	if err != nil {
		return nil, err
	}

	return &permiso, nil
}

func (r *permisoRepository) Update(id string, p *model.Permiso) error {

	permiso := model.Permiso{}
	permiso.ID = id

	err := r.db.Model(&permiso).Updates(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *permisoRepository) Delete(id string) error {

	// Delete Soft
	permiso := model.Permiso{}
	permiso.ID = id

	err := r.db.Delete(&permiso).Error

	if err != nil {
		return err
	}

	return nil
}
