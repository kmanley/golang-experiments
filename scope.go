package main

import "fmt"

func func1() int {
  x := 3
  return x
}

func main() {
  var z=func1()
  fmt.Println(z) //doit(x))
}
