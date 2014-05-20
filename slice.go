package main

import "fmt"
var print = fmt.Println

func main() {
  slice1 := []string {"string1", "string2", "string3", "string4"}
  print(slice1[2:3][0])
  print(slice1[3:])
  print(slice1[:3])
  print(slice1[:])
  //print(reversed(slice1[::-1])
}
