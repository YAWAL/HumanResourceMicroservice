package repository

import (
	"fmt"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"time"

	"github.com/YAWAL/ERP-common-lib/models"
	"github.com/google/uuid"

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
func (pg PostgresHRrepository) GetEmployees() []models.Employee {
	var Employees []models.Employee
	// SELECT * FROM employees
	rows := pg.conn.Table("employees")
	rows.Scan(&Employees)
	return Employees
}

//GetEmployeeById(id int64) *repo.Employee
func (pg PostgresHRrepository) GetEmployeeByID(id int64) *models.Employee {

	fmt.Println("GetEmployeeById not implemented yet")
	return nil
}

//GetEmployeeByName(name string) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByName(name string) []models.Employee {
	fmt.Println("GetEmployeeByName not implemented yet")
	return nil
}

//GetEmployeeByLastName(lastName string) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByLastName(lastName string) []*models.Employee {
	fmt.Println("GetEmployeeByLastName not implemented yet")
	return nil
}

//GetEmployeeByNameAndLastName(name, lastName string) []*repo.Employeeemployee_pk
func (pg PostgresHRrepository) GetEmployeeByNameAndLastName(name, lastName string) []*models.Employee {
	fmt.Println("GetEmployeeByNameAndLastName not implemented yet")
	return nil
}

//GetEmployeeByIdentificationNum(id int64) *repo.Employee
func (pg PostgresHRrepository) GetEmployeeByIdentificationNum(id int64) *models.Employee {
	fmt.Println("GetEmployeeByIdentificationNum not implemented yet")
	return nil
}

//GetEmployeeByPosition(position string) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByPosition(position string) []*models.Employee {
	fmt.Println("GetEmployeeByPosition not implemented yet")
	return nil
}

//GetEmployeeByJoinDate(joinDate time.Time) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByJoinDate(joinDate time.Time) []models.Employee {
	fmt.Println("GetEmployeeByJoinDate not implemented yet")
	return nil
}

//GetEmployeeByQuitDate(quitDate time.Time) []*repo.Employee
func (pg PostgresHRrepository) GetEmployeeByQuitDate(quitDate time.Time) []models.Employee {
	fmt.Println("GetEmployeeByQuitDate not implemented yet")
	return nil
}

//CreateEmployee(emp *repo.Employee) error
func (pg PostgresHRrepository) CreateEmployee(emp *models.Employee) error {
	//generate uuid
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	emp.ID = id
	pg.conn.Create(emp)
	info := emp.EmployeeInfo
	info.ID = id
	err = pg.CreateEmployeeInfo(&info)
	if err != nil {
		return err
	}
	address := emp.EmployeeInfo.EmployeeAddress
	address.ID = id
	err = pg.CreateEmployeeAddress(&address)
	if err != nil {
		return err
	}
	logging.Log.Infof("employee created: %s", emp)
	return nil
}

func (pg PostgresHRrepository) CreateEmployeeInfo(info *models.EmployeeInfo) error {
	return pg.conn.Create(info).Error
}

func (pg PostgresHRrepository) CreateEmployeeAddress(address *models.EmployeeAddress) error {
	return pg.conn.Create(address).Error
}

//UpdateEmployee(emp *repo.Employee) error
func (pg PostgresHRrepository) UpdateEmployee(emp *models.Employee) error {
	fmt.Println("UpdateEmployee not implemented yet")
	return nil
}

//DeleteEmployee(id int64) (bool, error)
func (pg PostgresHRrepository) DeleteEmployee(id int64) (bool, error) {
	fmt.Println("DeleteEmployee not implemented yet")
	return false, nil

}
