package main

import (
	"encoding/json"
	"github.com/pearleascent/volta"
)

func HelloWorld(ctx *volta.Ctx) error {
	return ctx.SendString("Hello World!")
}

func HelloWorldJson(ctx *volta.Ctx) error {
	return ctx.Redirect("/")
}

func main() {
	app := volta.New(volta.Config{
		Port:            "1337",
		JsonUnmarshaler: json.Unmarshal,
		JsonMarshaler:   json.Marshal,
	})

	app.Get("/", HelloWorld)
	app.Get("/json", HelloWorldJson)

	app.Run()
}
