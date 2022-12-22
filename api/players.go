package api

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/mbdeguzman/xmas_party/models"
	"github.com/uadmin/uadmin"
)

func PlayerDraw(w http.ResponseWriter, r *http.Request) {
	player_count, _ := strconv.ParseInt(r.FormValue("count"), 10, 64)
	employees := []models.Employees{}

	for i := range employees {
		uadmin.Preload(&employees[i])
	}
	players := []models.Employees{}

	for i := 0; i != int(player_count); i++ {
		uadmin.Filter(&employees, "registered = ? AND played = ?", true, false)
		rand.Seed(time.Now().Unix())
		min := 0
		max := len(employees)
		winner := employees[rand.Intn(max-min)+min]
		players = append(players, winner)
		win := models.Employees{}
		uadmin.Get(&win, "id=?", winner.ID)
		win.Played = true
		uadmin.Save(&win)
	}
	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"players": players,
	})
}
