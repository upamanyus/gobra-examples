// Gobra does not support this because "Function called from 'mayInit' context is not 'mayInit'."
package main

func f() int {
	return 0
}

var x int = f()

func main() {
}
