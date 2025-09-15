package main

import (
	"context"
	"log"
)

func main() {

	app := NewApp(context.Background())
	if err := app.Init(); err != nil {
		log.Fatal(err)
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
