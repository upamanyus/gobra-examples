// Gobra verifies this, even though it's a buggy program that panics.
package main

type foo struct {
	// Adding a field here causes Gobra verification on this file to fail.
	// a int
}

func main() {
	var s *foo
	var x foo = *s
	x = x
}
