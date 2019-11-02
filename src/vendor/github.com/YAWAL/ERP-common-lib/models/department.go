package models

// Department represents data about departments stored in database
type Department struct {
	DepartmentID   int
	DepartmentName string
	ChairmanID     int64
}
