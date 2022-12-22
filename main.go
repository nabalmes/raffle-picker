package main

import (
	"net/http"

	"github.com/mbdeguzman/xmas_party/api"
	"github.com/mbdeguzman/xmas_party/models"
	"github.com/mbdeguzman/xmas_party/views"
	"github.com/uadmin/uadmin"
)

func main() {
	DBConfig()
	RegisterModels()
	RegisterHandlers()
	SampleData()
	StartServer()
}

func DBConfig() {
	// uadmin.Database = &uadmin.DBSettings{
	// 	Type:     "mysql",
	// 	Name:     "xmas_party",
	// 	Host:     "localhost",
	// 	Password: "Allen is Great 200%",
	// 	Port:     3306,
	// }
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Name:     "xmas_party",
		Host:     "localhost",
		User:     "jmsemira",
		Password: "jmsemira",
		Port:     3306,
	}
	uadmin.Trail(uadmin.INFO, "Database Configured")
}

func RegisterModels() {
	uadmin.Trail(uadmin.INFO, "Register Models")
	uadmin.Register(
		models.Company{},
		models.Department{},
		models.Employees{},
	)
}

func RegisterHandlers() {
	uadmin.Trail(uadmin.INFO, "Register Handlers")
	http.HandleFunc("/index", uadmin.Handler(views.RaffleHandler))
	http.HandleFunc("/register", uadmin.Handler(views.RegisterHandler))
	http.HandleFunc("/raffle", uadmin.Handler(api.RaffleDraw))
	http.HandleFunc("/games", uadmin.Handler(views.GamesHandler))
	http.HandleFunc("/player", uadmin.Handler(api.PlayerDraw))
	http.HandleFunc("/register_emp", uadmin.Handler(api.Register))
	http.HandleFunc("/employee", uadmin.Handler(api.AllEmployees))
}

func StartServer() {
	uadmin.Port = 1118
	uadmin.StartServer()
}
