package api

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/mbdeguzman/xmas_party/models"
	"github.com/uadmin/uadmin"
)

func RaffleDraw(w http.ResponseWriter, r *http.Request) {
	employees := []models.Employees{}
	uadmin.Filter(&employees, "registered = ? AND raffle_qualified = ?", true, 1)

	for i := range employees {
		uadmin.Preload(&employees[i])
	}
	if len(employees) > 0 {
		rand.Seed(time.Now().Unix())
		min := 0
		max := len(employees)
		winner := employees[rand.Intn(max-min)+min]
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"winner": winner,
		})
		win := models.Employees{}
		uadmin.Get(&win, "id=?", winner.ID)
		win.RaffleQualified = false
		uadmin.Save(&win)
	} else {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"winner": "No qualified employees!",
		})
	}
}
