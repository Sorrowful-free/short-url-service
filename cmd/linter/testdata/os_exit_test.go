package main

import "os"

func testOsExit() {
	os.Exit(0) // want "calling os.Exit outside the main function of package main is not allowed"
}

func anotherFunction() {
	os.Exit(1) // want "calling os.Exit outside the main function of package main is not allowed"
}

func main() {
	os.Exit(0)
}
