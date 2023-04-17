package controllers

import (
	"github.com/kuropenguin/udemy-go-web-framework/framework/framework"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func GetStudent(ctx *framework.MyContext) {

	name := ctx.QueryKey("name", "default_name")
	studentResponse := StudentResponse{
		Name: name,
	}
	ctx.JSON(studentResponse)
}

func ListController(ctx *framework.MyContext) {
	ctx.WriteString("ListenController")
}

func ListItemController(ctx *framework.MyContext) {
	ctx.WriteString("list_item")
}

func UsersController(ctx *framework.MyContext) {
	ctx.WriteString("UsersController")
}
