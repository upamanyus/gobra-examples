// Causes gobra itself to crash
package main

func main() {
	m := make(map[any]bool)
	k := make([]int, 0)
	if m[k] {
	}
}
