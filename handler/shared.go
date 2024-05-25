package handler

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/goetian/gogen/models"
)

func render(w http.ResponseWriter, r *http.Request, com templ.Component) error {
	return com.Render(r.Context(), w)
}

func getAuthUser(r *http.Request) models.AuthUser {
	user, success := r.Context().Value(models.UserContextKey).(models.AuthUser)
	if !success {
		return models.AuthUser{}
	}
	return user
}

func MakeHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("Server Error", "err", err, "url", r.URL.Path)
		}
	}
}
