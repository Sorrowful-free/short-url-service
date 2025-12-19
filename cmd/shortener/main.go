package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var db *sql.DB //workaround for autotests

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	app := NewApp(ctx)
	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	errChan := make(chan error, 1)
	go func() {
		if err := app.Run(); err != nil {
			errChan <- err
		}
	}()

	select {
	case sig := <-sigChan:
		log.Printf("Received signal: %v. Starting graceful shutdown...", sig)
		if err := app.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down server: %v", err)
		} else {
			log.Println("Server shut down successfully")
		}
	case err := <-errChan:
		log.Fatal(err)
	}

	db.Close() //workaround for autotests
}
