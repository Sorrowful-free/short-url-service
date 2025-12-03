package other

func testPanic() {
	panic("test") // want "using panic is not allowed"
}
