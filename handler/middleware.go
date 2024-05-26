package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/goetian/gogen/models"
)

const test = false

func WithAuthentication(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.URL.Path, "/public") {

			next.ServeHTTP(w, r)
			return
		}
		user := models.AuthUser{}
		if test {
			user = models.AuthUser{
				Email:    "test@gmail.com",
				LoggedIn: true,
			}
		} else {
			user = models.AuthUser{}
		}
		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
