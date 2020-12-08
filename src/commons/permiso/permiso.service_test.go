package permiso

import (
	"fmt"
	"testing"
	"time"

	"github.com/nicobianchetti/Test-api2/database"
	"github.com/nicobianchetti/Test-api2/src/commons/model"
)

func TestFetchPermisoService(t *testing.T) {

	driver := database.Postgres
	database.New(driver)
	db := database.DB()

	rPermiso := NewPermisoRepository(db)
	sPermiso := NewPermisoService(rPermiso)

	testData := []struct {
		name string
		p    *model.Permiso
		err  string
	}{

		// {
		// 	name: "permiso found",
		// 	p:    permisoSampleS(),
		// },

		{
			name: "permiso not found",
			p:    &model.Permiso{ID: "123"},
			err:  "record not foundS",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			permiso, err := sPermiso.GetByID(tt.p.ID)

			fmt.Println("Permiso ", permiso)
			fmt.Println(err.Error())
			t.Fatalf(err.Error())

			// if err.Error() != tt.err {
			// 	t.Errorf("expected %v, got: %v", tt.err, err)
			// }

			if tt.err == "" {

				if permiso.ID == tt.p.ID {
					fmt.Println("Los ID son iguales")
				} else {
					fmt.Println("Los ID NO son iguales")
				}

				if *permiso != *tt.p {
					t.Fatalf("expected %v, permiso: %v", tt.p, permiso)
				}

			}

			// fmt.Println(err)

			// t.Errorf("expected %v, got: %v", tt.p, permiso)
		})
	}
}

func permisoSampleS() *model.Permiso {

	return &model.Permiso{
		ID:          "967d6b03-c41c-44a1-abce-25e9589293cf",
		Name:        "Permiso Post 2",
		Description: "Creación de permiso desde la api",
		Status:      true,
		Owner:       "Steve Vai",
		CreatedAt:   time.Date(2020, 10, 04, 18, 26, 8, 565766000, time.Now().Local().Location()),
		UpdatedAt:   time.Date(2020, 10, 04, 18, 26, 8, 565766000, time.Now().Local().Location()),
	}
}
