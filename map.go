package main

import "fmt"
var print = fmt.Println

func main() {
  people := make(map[string]int)
  people["Kevin"] = 46
  people["Magnus"] = 6
  people["Thor"] = 4
  people["Amy"] = 42
  print(people)
  //print(people.keys())
  age, ok := people["Kevin"]
  print(age, ok)
  age, ok = people["Fred"]
  print(age, ok)
  if age, ok = people["Magnus"]; ok {
     print("what the f?")
  }

  z := 3;4
  print(z)
}
