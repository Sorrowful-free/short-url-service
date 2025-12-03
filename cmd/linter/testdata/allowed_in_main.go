package main

import (
	"log"
	"os"
)

func main() {
	log.Fatal("this is allowed in main")
	os.Exit(0)
	panic("this is not allowed even in main") // want "using panic is not allowed"
}
