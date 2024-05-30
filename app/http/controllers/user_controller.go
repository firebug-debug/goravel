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

//func (r *UserController) Show(ctx http.Context) http.Response {
//	return ctx.Response().Success().Json(http.Json{
//		"Hello": "Goravel",
//	})
//}

func (r *UserController) Show(ctx http.Context) http.Response {
	var users []models.AdminUser
	err := facades.Orm().Query().Select("Id", "Name", "Username", "Balance", "Avatars").Get(&users)

	if err != nil {
		return ctx.Response().Success().Json(err)
	} else {
		for _, user := range users {
			fmt.Printf("user.name:%s", user.Name)
		}
		return ctx.Response().Success().Json(users)
	}
}

func (r *UserController) Add(ctx http.Context) http.Response {
	name := ctx.Request().Input("name", "goravel")
	username := ctx.Request().Input("username", "goravel")
	user := models.AdminUser{Name: name, Username: username, Password: ""}
	result := facades.Orm().Query().Create(&user)
	return ctx.Response().Success().Json(http.Json{"result": result})
}

func (r *UserController) DeleteUser(ctx http.Context) http.Response {
	var user models.AdminUser
	id := ctx.Request().Input("id", "0")
	if _, err := facades.Orm().Query().Delete(&user, id); err != nil {
		return ctx.Response().Success().Json(http.Json{"result": "err"})
	}

	return ctx.Response().Success().Json(http.Json{"result": "success"})

}

func (r *UserController) Login(ctx http.Context) http.Response {
	var user models.AdminUser
	err := facades.Orm().Query().Find(&user, 1)
	if err != nil {
		return ctx.Response().Success().Json(err)
	}
	token, err2 := facades.Auth().Guard("user").LoginUsingID(ctx, 1)
	if err2 != nil {
		return ctx.Response().Success().Json(err2)
	} else {
		_, err1 := facades.Auth().Guard("user").Parse(ctx, token)
		return ctx.Response().Success().Json(http.Json{
			"token": token,
			"err1":  err1,
		})
	}
}

func (r *UserController) UserInfo(ctx http.Context) http.Response {
	var user models.AdminUser
	//token := ctx.Request().Input("token", "")
	//if token != "" {
	//
	//}
	//err := facades.Auth().Parse(ctx, token)
	//token := ctx.Request().Header("Authorization", "")
	//facades.Auth().Guard("user").Parse(ctx, token)
	err := facades.Auth().Guard("user").User(ctx, &user)
	booll := errors.Is(err, auth.ErrorTokenExpired)
	bool1 := errors.Is(err, auth.ErrorInvalidKey)
	bool2 := errors.Is(err, auth.ErrorParseTokenFirst)
	return ctx.Response().Success().Json(http.Json{
		"user":  user,
		"err":   err,
		"booll": booll,
		"bool1": bool1,
		"bool2": bool2,
	})
}
