package main

import (
	"context"
	"database/sql"
	"log"
)

func main() {
	var db *sql.DB //workaraund for autotests
	app := NewApp(context.Background())
	if err := app.Init(); err != nil {
		log.Fatal(err)
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
	db.Close() //workaraund for autotests
}
