package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	uadmin.RenderHTML(w, r, "templates/register.html", data)
}
