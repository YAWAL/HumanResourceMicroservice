package testdata

import (
	"github.com/YAWAL/ERP-common-lib/models"
	"time"
)

var Employees = []models.Employee{
	models.Employee{
		EmployeeId:        1,
		Name:              "petro",
		LastName:          "1stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		AddressId:         1,
		BirthDate:         time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		EmployeeType:      "",
		Position:          "",
		JoinDate:          time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:          time.Now(),
		IsQuit:            false},
	models.Employee{
		EmployeeId:        2,
		Name:              "ivan",
		LastName:          "2stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		AddressId:         2,
		BirthDate:         time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		EmployeeType:      "",
		Position:          "",
		JoinDate:          time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:          time.Now(),
		IsQuit:            false},
	models.Employee{
		EmployeeId:        3,
		Name:              "3st",
		LastName:          "3stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		AddressId:         3, BirthDate: time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		EmployeeType: "",
		Position:     "",
		JoinDate:     time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:     time.Now(),
		IsQuit:       false},
	models.Employee{
		EmployeeId:        4,
		Name:              "ivan",
		LastName:          "3stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		AddressId:         4, BirthDate: time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		EmployeeType: "",
		Position:     "",
		JoinDate:     time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:     time.Now(),
		IsQuit:       false},
}
