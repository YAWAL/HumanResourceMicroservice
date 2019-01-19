package router

import (
	"github.com/YAWAL/HumanResourceMicroservice/src/HRusecases"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	hrApi := Router.PathPrefix("/HRMS").Subrouter()

	hrApi.HandleFunc("/", HRusecases.TempIndexPage)
	hrApi.HandleFunc("/employees", HRusecases.ShowAllEmployees).Methods("GET")
	hrApi.HandleFunc("/employees/{name:[a-z]+}", HRusecases.ShowAllEmployeesByName).Methods("GET")

	// profiling
	prof := Router.PathPrefix("/debug/pprof").Subrouter()
	prof.HandleFunc("/", pprof.Index)
	prof.HandleFunc("/cmdline", pprof.Cmdline)
	prof.HandleFunc("/symbol", pprof.Symbol)
	prof.HandleFunc("/trace", pprof.Trace)
	profile := Router.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("", pprof.Profile)
	profile.Handle("/goroutine", pprof.Handler("goroutine"))
	profile.Handle("/threadcreate", pprof.Handler("threadcreate"))
	profile.Handle("/heap", pprof.Handler("heap"))
	profile.Handle("/block", pprof.Handler("block"))
	profile.Handle("/mutex", pprof.Handler("mutex"))

}
