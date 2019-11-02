package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/YAWAL/ERP-common-lib/models"
	"github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
)

// TODO replace for static html
func TempIndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	t, err := template.ParseFiles(".\\templates\\index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = t.Execute(w, "Hello World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

}

// ShowAllEmployees renders all employees from database
func ShowAllEmployees(er database.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")

		employees := er.GetEmployees()
		if employees == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}


		renderedEmployees := models.PrepareRenderEmployee(employees)
		for _, value := range renderedEmployees {
			fmt.Println(value)

		}
		t, err := template.ParseFiles(".\\templates\\employees.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		err = t.Execute(w, renderedEmployees)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		return
	}
}

func ShowAllEmployeesByName(w http.ResponseWriter, r *http.Request) {

}

func CreateEmployee(er database.EmployeeRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(".\\templates\\createEmployee.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		//details := ContactDetails{
		//	Email:   r.FormValue("email"),
		//	Subject: r.FormValue("subject"),
		//	Message: r.FormValue("message"),
		//}
		employee :=models.Employee{
			Name:         r.FormValue("name"),
			LastName:     r.FormValue("lastName"),
			MiddleName:   "",
			Position:     r.FormValue("position"),
			IsQuit:       false,
			EmployeeInfo: models.EmployeeInfo{},
		}

		//defer r.Body.Close()
		//b, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	logging.Log.Errorf("CreateEmployee use case error: %s", err.Error())
		//	w.WriteHeader(http.StatusInternalServerError)
		//	w.Write([]byte(fmt.Sprint(err)))
		//	return
		//}
		//var employee models.Employee
		//if err = json.Unmarshal(b, &employee); err != nil {
		//	logging.Log.Errorf("CreateEmployee use case error: %s", err.Error())
		//	w.WriteHeader(http.StatusBadRequest)
		//	w.Write([]byte(fmt.Sprint(err)))
		//	return
		//}


		if err := er.CreateEmployee(&employee); err != nil {
			logging.Log.Errorf("CreateEmployee use case error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
}
