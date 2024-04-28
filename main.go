package main

import "main/app"

func main() {
	server, err := app.NewApp("8080")
	if err != nil {
		panic(err)
	}

	server.Start()
}
