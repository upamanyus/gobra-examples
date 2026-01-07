// Gobra does not support this because "got a that is not effective addressable"
package main

func main() {
	var a int
	x := &a
	x = x
}
