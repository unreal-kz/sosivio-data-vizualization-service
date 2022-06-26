package routes

import (
	"github.com/gorilla/mux"
	"github.com/unreal-kz/sosivio-data-vizualization-service/pkg/controllers"
)

var RegisterDataVisualRoutes = func(r *mux.Router) {
	r.HandleFunc("/data/", controllers.GetData).Methods("GET")
	r.HandleFunc("/data/", controllers.CreateData).Methods("POST")
	r.HandleFunc("/data/{id}", controllers.GetDataById).Methods("GET")

	// r.HandleFunc("/data/{id}", controllers.UpdateData).Methods("PUT")
	r.HandleFunc("/data/{id}", controllers.DeleteData).Methods("DELETE")
}
