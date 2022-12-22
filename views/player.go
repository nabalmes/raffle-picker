package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func GamesHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	uadmin.RenderHTML(w, r, "templates/player-selection.html", data)
}
