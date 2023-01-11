package main

import "github.com/pearleascent/volta"

func HelloWorld(ctx *volta.Ctx) error {
	return ctx.SendString("Hello World!")
}

func main() {
	app := volta.New(volta.Config{
		Port: "1337",
	})

	app.Get("/", HelloWorld)

	app.Run()
}
