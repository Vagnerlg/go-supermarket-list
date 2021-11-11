package main

import (
	"github.com/joho/godotenv"
	"github.com/vagnerlg/supermaketlist/src/config"
)

func main() {

	godotenv.Load()

	app := config.NewApp()

	routesItem := app.Http.Item
	routesItem.All("/")
	routesItem.FindById("/:id")
	routesItem.Insert("/")

	routeUser := app.Http.User
	routeUser.Insert("/")
	routeUser.Login("/login")

	app.Http.Drive.Run()
}
