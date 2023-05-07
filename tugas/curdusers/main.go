package main

import (
	"curdusers/configs"
	"curdusers/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8081"))
}
