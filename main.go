package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicobianchetti/Test-api2/src/commons/router"
)

const (
	port = "8080"
	ip   = ""
)

func main() {

	r := mux.NewRouter()
	router.SetupRoutesPermiso(r)

	addr := ip + ":" + port
	server := http.ListenAndServe(addr, r)
	log.Fatal(server)
}
