package main

import (
  "fmt"
  "runtime"
  "time"
)

var print = fmt.Println

var c chan int

func ready(w string, secs int) {
    time.Sleep(time.Duration(secs)*time.Second)
    print(w, "is ready")
    c <- 1
}

func test1() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  c = make(chan int)
  go ready("coffee", 2)
  go ready("tea", 1)
  <- c
  <- c
}


func main() {
    print(runtime.NumCPU())
    runtime.GOMAXPROCS(runtime.NumCPU())
    test1()
}

