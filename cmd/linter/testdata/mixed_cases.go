package main

import (
	"log"
	"os"
)

func helperFunction() {
	panic("not allowed")     // want "using panic is not allowed"
	log.Fatal("not allowed") // want "calling log.Fatal outside the main function of package main is not allowed"
	os.Exit(1)               // want "calling os.Exit outside the main function of package main is not allowed"
}

func init() {
	panic("not allowed in init")     // want "using panic is not allowed"
	log.Fatal("not allowed in init") // want "calling log.Fatal outside the main function of package main is not allowed"
	os.Exit(1)                       // want "calling os.Exit outside the main function of package main is not allowed"
}

func main() {
	log.Fatal("allowed")
	os.Exit(0)
}
