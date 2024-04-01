package view

import (
	"context"

	"github.com/goetian/gogen/models"
)

func DoAuthUser(ctx context.Context) models.AuthUser {
	user, success := ctx.Value(models.UserContextKey).(models.AuthUser)
	if !success {
		return models.AuthUser{}
	}
	return user
}
