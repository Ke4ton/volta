package main

import (
	"encoding/json"
	"fmt"
	"github.com/pearleascent/volta"
)

func HelloWorld(ctx *volta.Ctx) error {
	return ctx.SendString("Hello, " + ctx.Param("name", "") + "!")
}

func HelloWorldJson(ctx *volta.Ctx) error {
	fmt.Println(ctx.CurrentRoute())

	return ctx.SendJSON(volta.Map{
		"message": "Hello World!",
	})
}

func main() {
	app := volta.New(volta.Config{
		Port: "1337",
		// You have ability to select custom JSON marshaller and unmarshaler
		JsonUnmarshaler: json.Unmarshal,
		JsonMarshaler:   json.Marshal,
	})

	app.Use(func(ctx *volta.Ctx) error {
		if ctx.CurrentRoute() != "/json_test" {
			ctx.SendString("Hello, World!")
		}

		return ctx.Next()
	})

	app.Get("/hi/:name", HelloWorld)
	app.Get("/json_test", HelloWorldJson)

	app.Run()
}
