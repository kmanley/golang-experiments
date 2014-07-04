package main

import "fmt"

// tl;dr was wondering if you could return more than 2 values from a func--you can
func foo() (int, string, float32) {
	return 1, "hi", 1.23
}

func main() {
	x, y, z := foo()
	fmt.Println(x, y, z)

}
