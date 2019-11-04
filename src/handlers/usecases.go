package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/YAWAL/ERP-common-lib/models"
	"github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
)

// temporary HR index page
var HRwebpage = "<!DOCTYPE html><html>" +
	"<head>" +
	"<title>Human Resources Management</title>" +
	"</head>" +
	"<body>" +
	"<p>Human Resources service</p>" +
	"</body>" +
	"</html>"

func TempIndexPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, HRwebpage)
}

// ShowAllEmployees renders all employees from database
// GET/ employees
//func ShowAllEmployees(w http.ResponseWriter, r *http.Request) {
//
//}

func ShowAllEmployees(er database.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := er.GetEmployees()
		if err != nil {
			logging.Log.Errorf("ShowAllEmployees error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(employees)
		if err != nil {
			logging.Log.Errorf("ShowAllEmployees error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			logging.Log.Errorf("ShowAllEmployees error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
}

func ShowAllEmployeesByName(w http.ResponseWriter, r *http.Request) {

}

func CreateEmployee(er database.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logging.Log.Errorf("CreateEmployee error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		var employee models.Employee
		if err = json.Unmarshal(b, &employee); err != nil {
			logging.Log.Errorf("CreateEmployee error: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = er.CreateEmployee(&employee); err != nil {
			logging.Log.Errorf("CreateEmployee error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
}
