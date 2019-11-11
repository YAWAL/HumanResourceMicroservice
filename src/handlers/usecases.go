package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/YAWAL/ERP-common-lib/logging"
	"github.com/YAWAL/ERP-common-lib/models"
	"github.com/YAWAL/HumanResourceMicroservice/src/repository"

	"github.com/gorilla/mux"
)

const (
	positionParameter = "position"
	idParameter       = "id"
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
	fmt.Fprint(writer, HRwebpage)
}

// ShowAllEmployees retrieves all employees from MongoDB
// GET/ employees
func ShowAllEmployees(er repository.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := er.GetEmployees()
		if err != nil {
			logging.Log.Errorf("ShowAllEmployees error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logging.Log.Infof("ShowAllEmployees: %d employees have been retrieved", len(employees))
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
	}
}

// CreateEmployee creates employee document in MongoDB's database "hr", collection "employees"
// POST/ employees
func CreateEmployee(er repository.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
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
	}
}

// UpdateEmployee updates employee document in MongoDB's database "hr", collection "employees"
// PUT/ employees
func UpdateEmployee(er repository.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logging.Log.Errorf("UpdateEmployee error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		var employee models.Employee
		if err = json.Unmarshal(b, &employee); err != nil {
			logging.Log.Errorf("UpdateEmployee error: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = er.UpdateEmployee(&employee); err != nil {
			logging.Log.Errorf("UpdateEmployee error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// DeleteEmployee delets employee document from MongoDB's database "hr", collection "employees"
// DELETE/ employees/{id}
func DeleteEmployee(er repository.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		deleted, err := er.DeleteEmployee(mux.Vars(r)[idParameter])
		if err != nil {
			logging.Log.Errorf("DeleteEmployee error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if deleted {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

// ShowAllEmployeesByPosition retrieves all employees from MongoDB
// GET/ employees/{position}
func ShowAllEmployeesByPosition(er repository.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		position := mux.Vars(r)[positionParameter]
		employees, err := er.GetEmployeesByPosition(position)
		if err != nil {
			logging.Log.Errorf("ShowAllEmployeesByPosition error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logging.Log.Infof("ShowAllEmployeesByPosition: %d employees with position %s have been retrieved", len(employees), position)
		data, err := json.Marshal(employees)
		if err != nil {
			logging.Log.Errorf("ShowAllEmployeesByPosition error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			logging.Log.Errorf("ShowAllEmployeesByPosition error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
}
