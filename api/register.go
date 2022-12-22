package api

import (
	"net/http"
	"strconv"

	"github.com/mbdeguzman/xmas_party/models"
	"github.com/uadmin/uadmin"
)

func Register(w http.ResponseWriter, r *http.Request) {
	register, _ := strconv.ParseInt(r.FormValue("register"), 10, 64)
	employees := models.Employees{}
	uadmin.Get(&employees, "employee_number", register)

	employees.Registered = true
	uadmin.Save(&employees)
}

func AllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := []models.Employees{}
	uadmin.Filter(&employees, "registered", 0)

	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"employee": employees,
	})

}
