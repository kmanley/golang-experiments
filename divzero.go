package main

import "fmt"

func doit(x int) int {
	return x / 0.0
}

func main() {
	var x = 100
	fmt.Println(doit(x))
}
