// Gobra does not support this because "array arr is not addressable"

package main

func main() {
	var arr [10]int
	sl := arr[:]
	sl = sl
}
