// Gobra verifies this, even though it's a buggy program that panics.
package main

type foo struct {
	x int
}

func main() {
	var s *foo
	var x *int = &s.x
	x = x

	// This actually cause a gobra verification failure:
	// var y int = s.x
	// y = y
}
