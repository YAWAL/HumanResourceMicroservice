package router

import (
	"net/http"

	"github.com/YAWAL/HumanResourceMicroservice/src/handlers"

	"github.com/gorilla/mux"
)

func InitRouter() (r *mux.Router) {
	r = mux.NewRouter()
	api := r.PathPrefix("/humanResources").Subrouter()
	api.HandleFunc("/", handlers.TempIndexPage)
	api.HandleFunc("/employees", handlers.ShowAllEmployees).Methods(http.MethodGet)
	api.HandleFunc("/employees/{name:[a-z]+}", handlers.ShowAllEmployeesByName).Methods(http.MethodGet)
	return r
}
