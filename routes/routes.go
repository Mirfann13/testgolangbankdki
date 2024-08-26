package routes

import (
	"inventory-api/controllers"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/stocks", controllers.CreateStock).Methods("POST")
	r.HandleFunc("/stocks", controllers.ListStock).Methods("GET")
	r.HandleFunc("/stocks/{id:[0-9]+}", controllers.GetStock).Methods("GET")
	r.HandleFunc("/stocks/{id:[0-9]+}", controllers.UpdateStock).Methods("PUT")
	r.HandleFunc("/stocks/{id:[0-9]+}", controllers.DeleteStock).Methods("DELETE")
	return r
}
