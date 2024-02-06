package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
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
