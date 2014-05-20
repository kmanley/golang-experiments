package main

import "fmt"

func doloop() {
  for i:=0; i<10; i++ {
    if i % 2 == 0 {
      fmt.Println(i, "is even")
    } else {
      fmt.Println(i, "is odd")
    }
  }
}

func main() {
  doloop()
}
