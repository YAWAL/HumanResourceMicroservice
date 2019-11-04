package router

import (
	"net/http"

	"github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/YAWAL/HumanResourceMicroservice/src/handlers"

	"github.com/gorilla/mux"
)

func InitRouter(er database.EmployeeRepository) (r *mux.Router) {
	r = mux.NewRouter()
	api := r.PathPrefix("/humanResources").Subrouter()
	api.HandleFunc("/", handlers.TempIndexPage)
	api.HandleFunc("/employees", handlers.ShowAllEmployees(er)).Methods(http.MethodGet)
	api.HandleFunc("/employees", handlers.CreateEmployee(er)).Methods(http.MethodPost)
	api.HandleFunc("/employees/{name:[a-z]+}", handlers.ShowAllEmployeesByName).Methods(http.MethodGet)
	return r
}
