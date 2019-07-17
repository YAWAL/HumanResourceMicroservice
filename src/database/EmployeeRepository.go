package database

import (
	"github.com/YAWAL/ERP-common-lib/models"
)

// where this var should be
type Employee models.Employee

type EmployeeRepository interface {
	CreateEmployee(emp *Employee) error
	UpdateEmployee(emp *Employee) error
	DeleteEmployee(id int64) (bool, error)
	GetEmployees() []*Employee

	//GetEmployeeById(id int64) *Employee
	//GetEmployeeByName(name string) []*Employee
	//GetEmployeeByLastName(lastName string) []*Employee
	//GetEmployeeByNameAndLastName(name, lastName string) []*Employee
	//GetEmployeeByIdentificationNum(id int64) *Employee
	//GetEmployeeByPosition(position string) []*Employee
	//GetEmployeeByJoinDate(joinDate time.Time) []*Employee
	//GetEmployeeByQuitDate(quitDate time.Time) []*Employee
}
