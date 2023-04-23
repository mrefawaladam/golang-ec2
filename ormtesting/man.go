package main

import (
	"orm/config"
	"orm/route"
)

func main() {
	config.InitDB()

	e := route.New()

	e.Logger.Fatal(e.Start(":8000"))
}
