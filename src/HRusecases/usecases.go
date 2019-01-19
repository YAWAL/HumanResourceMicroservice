package HRusecases

import (
	"fmt"
	"net/http"

	"github.com/YAWAL/ERP/back-end/HumanResourcesService/HRrepository"
	"github.com/YAWAL/ERP/back-end/common"
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
func ShowAllEmployees(w http.ResponseWriter, r *http.Request) {
	common.RenderJSON(w, (HRrepository.Employee{}).GetEmployees())
}

func ShowAllEmployeesByName(w http.ResponseWriter, r *http.Request) {
	common.RenderJSON(w, (HRrepository.Employee{}).GetEmployeeByName)
}
