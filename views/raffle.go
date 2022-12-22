package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func RaffleHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	uadmin.RenderHTML(w, r, "templates/index.html", data)
}
