package main

import "fmt"

func main() {
	var s1 string
	s2 := ""
	if s1 == s2 {
		fmt.Println("'' is the zero value of string")
	} else {
		fmt.Println(s1, "is the zero value of string")
	}
}
