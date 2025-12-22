package main

import (
	"context"
	"database/sql"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	var db *sql.DB //workaround for autotests

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, cancel = signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	app := NewApp(ctx)
	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	db.Close() //workaround for autotests
}
