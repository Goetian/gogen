package handler

import (
	"net/http"

	"github.com/goetian/gogen/view/home"
)

func HandlerLandingPage(w http.ResponseWriter, r *http.Request) error {

	return render(w, r, home.Index())
}
