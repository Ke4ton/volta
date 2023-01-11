package main

import (
	"github.com/pearleascent/volta"
)

func HelloWorld(ctx *volta.Ctx) error {
	return ctx.Status(volta.StatusOK).SendString("Hello World!")
}

func HelloWorldJson(ctx *volta.Ctx) error {
	return ctx.SendJSON(volta.Map{
		"message": "Hello World!",
	})
}

func main() {
	app := volta.New(volta.Config{
		Port: "1337",
		// You have ability to select custom JSON marshaller and unmarshaler
		//JsonUnmarshaler: json.Unmarshal,
		//JsonMarshaler:   json.Marshal,
	})

	app.Get("/", HelloWorld)
	app.Get("/json_test", HelloWorldJson)

	app.Run()
}
