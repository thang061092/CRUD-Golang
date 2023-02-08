package controllers

import (
	"api/helper"
	"api/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) FindUserById() {
	id, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))

	user := models.GetUserById(id)
	helper.SendResponse(&u.Controller, helper.HTTP_OK, helper.SUCCESS, user)
}

func (u *UserController) GetPaginateUser() {
	users := models.GetPaginateUser()
	helper.SendResponse(&u.Controller, helper.HTTP_OK, helper.SUCCESS, users)
}

func (u *UserController) UpdateUser() {
	id, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))

	var data models.User
	if err := u.ParseForm(&data); err != nil {
		helper.SendResponse(&u.Controller, helper.HTTP_BAD_REQUEST, helper.FAIL, err.Error())
		return
	}
	user := models.UpdateUser(id, &data)
	helper.SendResponse(&u.Controller, helper.HTTP_OK, helper.SUCCESS, user)
}

func (u *UserController) CreateUser() {
	var data models.User
	if err := u.ParseForm(&data); err != nil {
		helper.SendResponse(&u.Controller, helper.HTTP_BAD_REQUEST, helper.FAIL, err.Error())
		return
	}
	id, err := models.CreateUser(&data)
	if err == nil {
		helper.SendResponse(&u.Controller, helper.HTTP_OK, helper.SUCCESS, id)
	} else {
		helper.SendResponse(&u.Controller, helper.HTTP_OK, err.Error(), nil)
	}
}
