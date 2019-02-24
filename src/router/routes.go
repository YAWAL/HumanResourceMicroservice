package router

import (
	"github.com/YAWAL/HumanResourceMicroservice/src/handlers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	var router *mux.Router

	router = mux.NewRouter()

	hrApi := router.PathPrefix("/HRMS").Subrouter()

	hrApi.HandleFunc("/", handlers.TempIndexPage)
	hrApi.HandleFunc("/employees", handlers.ShowAllEmployees).Methods("GET")
	hrApi.HandleFunc("/employees/{name:[a-z]+}", handlers.ShowAllEmployeesByName).Methods("GET")

	return router
}
