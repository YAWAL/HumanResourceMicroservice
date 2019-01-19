package database

import (
	"time"

	"github.com/YAWAL/ERP-common-lib/models"
)

type EmployeeRepository interface {
	CreateEmployee(emp *models.Employee) error
	UpdateEmployee(emp *models.Employee) error
	DeleteEmployee(id int64) (bool, error)

	GetEmployees() []*models.Employee
	GetEmployeeById(id int64) *models.Employee
	GetEmployeeByName(name string) []*models.Employee
	GetEmployeeByLastName(lastName string) []*models.Employee
	GetEmployeeByNameAndLastName(name, lastName string) []*models.Employee
	GetEmployeeByIdentificationNum(id int64) *models.Employee
	GetEmployeeByPosition(position string) []*models.Employee
	GetEmployeeByJoinDate(joinDate time.Time) []*models.Employee
	GetEmployeeByQuitDate(quitDate time.Time) []*models.Employee
}

type AddressRepository interface {
	GetEmployeeAddress(employeeId int64) *models.EmployeeAddress
}
