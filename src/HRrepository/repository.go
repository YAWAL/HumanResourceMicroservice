package HRrepository

import (
	"fmt"
	"github.com/YAWAL/ERP/back-end/HumanResourcesService/model"
	"github.com/YAWAL/ERP/back-end/database"
	"github.com/YAWAL/ERP/back-end/tool/logging"
	"time"
)

// should implements EmployeeRepository interface
type Employee model.Employee

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

func (e Employee) GetEmployeeById(id int64) *model.Employee {

	fmt.Println("GetEmployeeById not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByName(name string) []*model.Employee {
	fmt.Println("GetEmployeeByName not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByLastName(lastName string) []*model.Employee {
	fmt.Println("GetEmployeeByLastName not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByNameAndLastName(name, lastName string) []*model.Employee {
	fmt.Println("GetEmployeeByNameAndLastName not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByIdentificationNum(id int64) *model.Employee {
	fmt.Println("GetEmployeeByIdentificationNum not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByPosition(position string) []*model.Employee {
	fmt.Println("GetEmployeeByPosition not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByJoinDate(joinDate time.Time) []*model.Employee {
	fmt.Println("GetEmployeeByJoinDate not implemented yet")
	return nil
}

func (e Employee) GetEmployeeByQuitDate(quitDate time.Time) []*model.Employee {
	fmt.Println("GetEmployeeByQuitDate not implemented yet")
	return nil
}

func CreateEmployee(emp *model.Employee) error {
	fmt.Println("CreateEmployee not implemented yet")

	return nil
}

func UpdateEmployee(emp *model.Employee) error {
	fmt.Println("UpdateEmployee not implemented yet")
	return nil
}

func DeleteEmployee(id int64) (bool, error) {
	fmt.Println("DeleteEmployee not implemented yet")
	return false, nil

}
