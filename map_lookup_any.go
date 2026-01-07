// Causes gobra itself to crash; perhaps because of the implicit conversion?
package main

func main() {
	m := make(map[any]bool)
	var k string
	if m[k] {
	}
}
