package main

import "log"

func testLogFatal() {
	log.Fatal("error") // want "calling log.Fatal outside the main function of package main is not allowed"
}

func anotherFunction() {
	log.Fatal("fatal error") // want "calling log.Fatal outside the main function of package main is not allowed"
}

func main() {
	log.Fatal("this is allowed")
}

