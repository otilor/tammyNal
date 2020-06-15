package tammyNal

import (
	"net/http"
)

func Index (w http.ResponseWriter, r *http.Request) {
	render(w, "tammyNal.png", r)
}