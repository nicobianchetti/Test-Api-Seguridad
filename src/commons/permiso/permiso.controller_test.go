package permiso

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// const (
// 	port = "8080"
// 	ip   = ""
// )

// var roPermiso permisoRouter
// var r *mux.Router

// func TestMain(m *testing.M) {
// 	r = mux.NewRouter()

// 	driver := database.Postgres
// 	database.New(driver)
// 	db := database.DB()

// 	rPermiso := NewPermisoRepository(db)
// 	sPermiso := NewPermisoService(rPermiso)
// 	cPermiso := NewPermisoController(sPermiso)
// 	roPermiso = NewPermisoRouterTest(r.PathPrefix("/v1/permiso").Subrouter(), cPermiso)

// 	code := m.Run()
// 	os.Exit(code)
// }

func TestFetchPermisoController(t *testing.T) {
	testData := []struct {
		name   string
		p      *DTOPermisoResponse
		status int
		err    string
	}{

		{
			name:   "permiso found",
			p:      permisoSampleC(),
			status: http.StatusOK,
		},

		{
			name:   "permiso not found",
			p:      &DTOPermisoResponse{ID: "123"},
			status: http.StatusNotFound,
			err:    "Permiso Not found",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			uri := fmt.Sprintf("/v1/permiso/permiso/%s", tt.p.ID)
			req, err := http.NewRequest(http.MethodGet, uri, nil)

			if err != nil {
				t.Fatalf("could not created request: %v", err)
			}

			rr := executeRequest(req)

			res := rr.Result()

			defer res.Body.Close()

			if tt.status != res.StatusCode {
				t.Errorf("expected %d, got: %d", tt.status, res.StatusCode)
			}

			b, err := ioutil.ReadAll(res.Body)

			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tt.err == "" {
				var got *DTOPermisoResponse
				err = json.Unmarshal(b, &got)
				if err != nil {
					t.Fatalf("could not unmarshall response %v", err)
				}

				if *got != *tt.p {
					t.Fatalf("expected %v, got: %v", tt.p, got)
				}
			}
		})
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	return rr
}

func permisoSampleC() *DTOPermisoResponse {

	return &DTOPermisoResponse{
		ID:          "967d6b03-c41c-44a1-abce-25e9589293cf",
		Name:        "Permiso Post 2",
		Description: "Creación de permiso desde la api",
		Status:      true,
		Owner:       "Steve Vai",
		CreatedAt:   time.Date(2020, 10, 04, 18, 26, 8, 565766000, time.Now().Local().Location()),
		UpdatedAt:   time.Date(2020, 10, 04, 18, 26, 8, 565766000, time.Now().Local().Location()),
	}
}
