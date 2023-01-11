package main

import "github.com/pearleascent/volta"

func main() {
	app := volta.New(volta.Config{
		Port: "1337",
	})

	app.Run()
}
