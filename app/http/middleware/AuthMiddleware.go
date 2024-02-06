package middleware

import (
	"errors"
	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AuthMiddleware() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization", "")
		err := facades.Auth().Guard("user").Parse(ctx, token)
		tokenExpired := errors.Is(err, auth.ErrorTokenExpired)
		if tokenExpired {
			errors.New("token expired")
		}
		invalidKey := errors.Is(err, auth.ErrorInvalidKey)
		if invalidKey {
			errors.New("invalid key")
		}
		ctx.Request().Next()
	}
}
