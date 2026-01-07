// Gobra verifies this, even though it's a buggy program that panics.
package main

type A struct {
	a any
}

func main() {
	m := make(map[A]bool)
	k := A{a: []int(nil)}
	if m[k] {
	}
}
