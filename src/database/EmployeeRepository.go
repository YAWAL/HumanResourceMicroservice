package database

import (
	"github.com/YAWAL/ERP-common-lib/models"
)

// EmployeeRepository is a methods set for datastores for employees related information
type EmployeeRepository interface {
	CreateEmployee(emp *models.Employee) error
	UpdateEmployee(emp *models.Employee) error
	DeleteEmployee(id string) (bool, error)
	GetEmployees() (emps []*models.Employee, err error)
	GetEmployeesByPosition(position string) (emps []*models.Employee, err error)

	//GetEmployeeById(id int64) *Employee
	//GetEmployeeByName(name string) []*Employee
	//GetEmployeeByLastName(lastName string) []*Employee
	//GetEmployeeByNameAndLastName(name, lastName string) []*Employee
	//GetEmployeeByIdentificationNum(id int64) *Employee
	//GetEmployeeByJoinDate(joinDate time.Time) []*Employee
	//GetEmployeeByQuitDate(quitDate time.Time) []*Employee
}
