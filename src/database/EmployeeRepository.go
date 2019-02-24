package database

import (
	repo "github.com/YAWAL/HumanResourceMicroservice/src/HRrepository"
	"time"
)

type EmployeeRepository interface {
	CreateEmployee(emp *repo.Employee) error
	UpdateEmployee(emp *repo.Employee) error
	DeleteEmployee(id int64) (bool, error)

	GetEmployees() []*repo.Employee
	GetEmployeeById(id int64) *repo.Employee
	GetEmployeeByName(name string) []*repo.Employee
	GetEmployeeByLastName(lastName string) []*repo.Employee
	GetEmployeeByNameAndLastName(name, lastName string) []*repo.Employee
	GetEmployeeByIdentificationNum(id int64) *repo.Employee
	GetEmployeeByPosition(position string) []*repo.Employee
	GetEmployeeByJoinDate(joinDate time.Time) []*repo.Employee
	GetEmployeeByQuitDate(quitDate time.Time) []*repo.Employee
}
