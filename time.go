package main

import (
  "fmt"
  "time"
)
var print = fmt.Println

func testAfterFunc() {
  x := 3
  y := 4
  adder := func() {
      print("x*y=", x*y)
  }
  d,_ := time.ParseDuration("5s")
  time.AfterFunc(d, adder)
}

func main() {
  start := time.Now()
  elapsed := time.Since(start)
  print(elapsed)
  
  testAfterFunc()
  print("sleeping")
  d, _ := time.ParseDuration("2s")
  print("duration is", d)
  print(time.Duration(1e6)) 
  print(time.ParseDuration("1m45s")) // returns duration, error
  
  time.Sleep(d)
  print("done!")
  
}