// Gobra verifies this, even though it's a buggy program that panics.
package main

type Any struct {
	a any
}

func main() {
	var x, y Any
	x.a = Any{a: make(map[int]int)}
	y.a = Any{a: make(map[int]int)}
	if x == y {
	}
}
