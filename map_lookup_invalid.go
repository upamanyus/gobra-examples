// Gobra verifies this, even though it's a buggy program that panics.
package main

type foo struct {
	a any
}

func main() {
	m := make(map[foo]bool)
	f := foo{a: []int(nil)}
	if m[f] {
	}
}
