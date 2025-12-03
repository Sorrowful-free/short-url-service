package main

import "os"

func testExit() {
	os.Exit(0)
}

func anotherFunction() {
	testExit()
}

func main() {
	testExit()
	anotherFunction()
}
