package controller

import (
	"net/http"

	"binus.ac.id/model"
	"binus.ac.id/view"
)

func EchoController(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("v")
	data := model.EchoData{Value: value}

	view.EchoTemplate().Execute(w, data)
}
