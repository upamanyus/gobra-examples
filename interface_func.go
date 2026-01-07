// Unsupported by gobra because "The type func() () is not supported for interface"

package main

func main() {
	var a any
	a = func(){}
	a = a
}
