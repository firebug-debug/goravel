package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
	"goravel/app/http/controllers"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})
	//.Middleware(middleware.Cors())
	userController := controllers.NewUserController()
	facades.Route.Get("/users/{id}", userController.Show)
	facades.Route.Get("/usersAdd", userController.Add)
	facades.Route.Get("/user/delete", userController.DeleteUser)
	facades.Route.Get("/user/Login", userController.Login)
	facades.Route.Get("/user/info", userController.UserInfo)
}
