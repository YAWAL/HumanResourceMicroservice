package repository

import (
	"fmt"
	"time"

	db "github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/jinzhu/gorm"
)

// should implements EmployeeRepository interface
func NewPostgresRepository(conn *gorm.DB) PostgresHRrepository {
	return PostgresHRrepository{conn: conn}
}

type PostgresHRrepository struct {
	//db.EmployeeRepository
	conn *gorm.DB
}

//GetEmployees() []*repo.Employee
func (pg PostgresHRrepository) GetEmployees() *[]db.Employee {
	var Employees []db.Employee
	// SELECT * FROM employees
	rows := pg.conn.Table("employees")
	rows.Scan(Employees)

	return &Employees

}

//GetEmployeeById(id int64) *repo.Employee
func (pg PostgresHRrepository) GetEmployeeByID(id int64) *db.Employee {

	fmt.Println("GetEmployeeById not implemented yet")
	return nil
}

//GetEmployeeByName(name string) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByName(name string) []db.Employee {
	fmt.Println("GetEmployeeByName not implemented yet")
	return nil
}

//GetEmployeeByLastName(lastName string) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByLastName(lastName string) []*db.Employee {
	fmt.Println("GetEmployeeByLastName not implemented yet")
	return nil
}

//GetEmployeeByNameAndLastName(name, lastName string) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByNameAndLastName(name, lastName string) []*db.Employee {
	fmt.Println("GetEmployeeByNameAndLastName not implemented yet")
	return nil
}

//GetEmployeeByIdentificationNum(id int64) *repo.Employee
func (pg PostgresHRrepository) GetEmployeeByIdentificationNum(id int64) *db.Employee {
	fmt.Println("GetEmployeeByIdentificationNum not implemented yet")
	return nil
}

//GetEmployeeByPosition(position string) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByPosition(position string) []*db.Employee {
	fmt.Println("GetEmployeeByPosition not implemented yet")
	return nil
}

//GetEmployeeByJoinDate(joinDate time.Time) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByJoinDate(joinDate time.Time) []db.Employee {
	fmt.Println("GetEmployeeByJoinDate not implemented yet")
	return nil
}

//GetEmployeeByQuitDate(quitDate time.Time) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByQuitDate(quitDate time.Time) []db.Employee {
	fmt.Println("GetEmployeeByQuitDate not implemented yet")
	return nil
}

//CreateEmployee(emp *repo.Employee) error
func (pg PostgresHRrepository) CreateEmployee(emp *db.Employee) error {
	fmt.Println("CreateEmployee not implemented yet")

	return nil
}

//UpdateEmployee(emp *repo.Employee) error
func (pg PostgresHRrepository) UpdateEmployee(emp *db.Employee) error {
	fmt.Println("UpdateEmployee not implemented yet")
	return nil
}

//DeleteEmployee(id int64) (bool, error)
func (pg PostgresHRrepository) DeleteEmployee(id int64) (bool, error) {
	fmt.Println("DeleteEmployee not implemented yet")
	return false, nil

}
