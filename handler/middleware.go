package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/goetian/gogen/models"
)

func WithAuthentication(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {

			next.ServeHTTP(w, r)
			return
		}
		user := models.AuthUser{}
		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
