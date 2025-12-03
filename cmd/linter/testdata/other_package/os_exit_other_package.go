package other

import "os"

func testOsExit() {
	os.Exit(0) // want "calling os.Exit outside the main function of package main is not allowed"
}
