package controller

import (
	"net/http"

	"binus.ac.id/view"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	view.HomeTemplate().Execute(w, nil)
}
