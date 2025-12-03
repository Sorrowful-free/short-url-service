package other

import "log"

func testLogFatal() {
	log.Fatal("error") // want "calling log.Fatal outside the main function of package main is not allowed"
}

