package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/unreal-kz/sosivio-data-vizualization-service/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterDataVisualRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
