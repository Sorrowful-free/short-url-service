package main

func testPanic() {
	panic("test") // want "using panic is not allowed"
}

func anotherFunction() {
	panic("error") // want "using panic is not allowed"
}

func main() {
	panic("main panic") // want "using panic is not allowed"
}

