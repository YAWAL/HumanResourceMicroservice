package database

import (
	"github.com/YAWAL/ERP-common-lib/models"
)

type EmployeeAddress models.EmployeeAddress

type AddressRepository interface {
	GetEmployeeAddress(employeeID int64) *EmployeeAddress
}
