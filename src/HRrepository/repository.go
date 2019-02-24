package HRrepository

import (
	"fmt"
	"github.com/YAWAL/ERP-common-lib/models"
	"github.com/YAWAL/HumanResourceMicroservice/src/database"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
	"time"
)

// should implements EmployeeRepository interface
type Employee models.Employee

func (e Employee) GetEmployees() []Employee {
	db, err := database.PGconn()
	if err != nil {
		logging.Log.Error("DB error has occurred: ", err)
		return nil
	}

	var Employees []Employee
	// SELECT * FROM employees
	rows := db.Table("employees")
	rows.Scan(&Employees)

	return Employees

}

func (e Employee) GetEmployeeByID(id int64) *Employee {

	fmt.Println("GetEmployeeById not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByName(name string) []Employee {
	fmt.Println("GetEmployeeByName not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByLastName(lastName string) []Employee {
	fmt.Println("GetEmployeeByLastName not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByNameAndLastName(name, lastName string) []Employee {
	fmt.Println("GetEmployeeByNameAndLastName not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByIdentificationNum(id int64) *Employee {
	fmt.Println("GetEmployeeByIdentificationNum not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByPosition(position string) []Employee {
	fmt.Println("GetEmployeeByPosition not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByJoinDate(joinDate time.Time) []Employee {
	fmt.Println("GetEmployeeByJoinDate not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByQuitDate(quitDate time.Time) []Employee {
	fmt.Println("GetEmployeeByQuitDate not implemented yet")
	return nil
}

func CreateEmployee(emp Employee) error {
	fmt.Println("CreateEmployee not implemented yet")

	return nil
}

func UpdateEmployee(emp Employee) error {
	fmt.Println("UpdateEmployee not implemented yet")
	return nil
}

func DeleteEmployee(id int64) (bool, error) {
	fmt.Println("DeleteEmployee not implemented yet")
	return false, nil

}
