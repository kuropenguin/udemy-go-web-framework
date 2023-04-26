package controllers

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"

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

func ListItemPictureController(ctx *framework.MyContext) {
	listID := ctx.GetParam(":list_id", "")
	pictureID := ctx.GetParam(":picture_id", "")

	type OUTPUT struct {
		ListID    string `json:"list_id"`
		PictureID string `json:"picture_id"`
	}
	ctx.JSON(&OUTPUT{
		ListID:    listID,
		PictureID: pictureID,
	})
}

func PostPageController(ctx *framework.MyContext) {
	ctx.WriteString(`
		<html>
			<body>
				<form action="/posts" method="post" enctype="multipart/form-data">
					<input type="text" name="title" />
					<input type="text" name="name" />
					<input type="file" name="file" />
					<input type="submit" />
					</form>
					</body>
					</html>
	`)
}

func PostController(ctx *framework.MyContext) {
	name := ctx.FormKey("name", "default_name")
	age := ctx.FormKey("age", "20")
	fileInfo, err := ctx.FormFile("file")

	if err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s_%s_%s", name, age, fileInfo.Filename), fileInfo.Data, fs.ModePerm)
	if err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
	}

	ctx.WriteString("post")
}
