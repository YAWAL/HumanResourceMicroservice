package router

import (
	"net/http"

	"github.com/YAWAL/HumanResourceMicroservice/src/handlers"
	"github.com/YAWAL/HumanResourceMicroservice/src/repository"

	"github.com/gorilla/mux"
)

// InitRouter creates gorilla mux router and mapped paths with handler functions
func InitRouter(er repository.EmployeeRepository) (r *mux.Router) {
	r = mux.NewRouter()
	api := r.PathPrefix("/humanResources").Subrouter()
	api.HandleFunc("/", handlers.TempIndexPage)
	api.HandleFunc("/employees", handlers.ShowAllEmployees(er)).Methods(http.MethodGet)
	api.HandleFunc("/employees", handlers.CreateEmployee(er)).Methods(http.MethodPost)
	api.HandleFunc("/employees", handlers.UpdateEmployee(er)).Methods(http.MethodPut)
	api.HandleFunc("/employees/{id:[0-9a-zA-Z]+}", handlers.DeleteEmployee(er)).Methods(http.MethodDelete)
	api.HandleFunc("/employees/{position:[0-9a-zA-Z]+}", handlers.ShowAllEmployeesByPosition(er)).Methods(http.MethodGet)
	return r
}
