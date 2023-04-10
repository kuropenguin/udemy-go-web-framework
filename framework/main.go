package main

import (
	"github.com/kuropenguin/udemy-go-web-framework/framework/controllers"
	"github.com/kuropenguin/udemy-go-web-framework/framework/framework"
)

func main() {
	engin := framework.NewEngine()
	router := engin.Router
	router.Get("/students", controllers.GetStudent)
	router.Get("/lists", controllers.ListenController)
	router.Get("/users", controllers.UsersController)
	engin.Run()
}
