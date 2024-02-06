package controllers

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {

	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) {
	var users []models.AdminUser
	err := facades.Orm.Query().Select("Id", "Name", "Username", "Balance", "Avatars").Get(&users)

	if err != nil {
		ctx.Response().Success().Json(err)
	} else {
		for _, user := range users {
			fmt.Printf("user.name:%s", user.Name)
		}
		ctx.Response().Success().Json(users)
	}
}

func (r *UserController) Add(ctx http.Context) {
	name := ctx.Request().Form("name", "goravel")
	username := ctx.Request().Form("username", "goravel")
	user := models.AdminUser{Name: name, Username: username, Password: ""}
	result := facades.Orm.Query().Create(&user)
	ctx.Response().Success().Json(http.Json{"result": result})
}

func (r *UserController) DeleteUser(ctx http.Context) {
	var user models.AdminUser
	id := ctx.Request().Form("id", "0")
	facades.Orm.Query().Delete(&user, id)

	ctx.Response().Success().Json(http.Json{"result": "success"})

}

func (r *UserController) Login(ctx http.Context) {
	var user models.AdminUser
	facades.Orm.Query().Find(&user, 1)

	token, err := facades.Auth.Guard("user").LoginUsingID(ctx, 1)
	if err != nil {
		ctx.Response().Success().Json(err)
	} else {
		err1 := facades.Auth.Guard("user").Parse(ctx, token)
		ctx.Response().Success().Json(http.Json{
			"token": token,
			"err1":  err1,
		})
	}
}

func (r *UserController) UserInfo(ctx http.Context) {
	var user models.AdminUser
	//token := ctx.Request().Form("token", "")
	//if token != "" {
	//
	//}
	//err := facades.Auth.Parse(ctx, token)
	//token := ctx.Request().Header("Authorization", "")
	//facades.Auth.Guard("user").Parse(ctx, token)
	err := facades.Auth.Guard("user").User(ctx, &user)
	booll := errors.Is(err, auth.ErrorTokenExpired)
	bool1 := errors.Is(err, auth.ErrorInvalidKey)
	bool2 := errors.Is(err, auth.ErrorParseTokenFirst)
	ctx.Response().Success().Json(http.Json{
		"user":  user,
		"err":   err,
		"booll": booll,
		"bool1": bool1,
		"bool2": bool2,
	})
}
