package models

import "github.com/uadmin/uadmin"

type Employees struct {
	uadmin.Model
	Name            string `uadmin:"read_only"`
	EmployeeNumber  string `uadmin:"read_only"`
	Department      Department
	DepartmentID    uint
	Company         Company
	CompanyID       uint
	Registered      bool
	RaffleQualified bool
	Played          bool `uadmin:"default_value:false"`
}
