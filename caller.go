package main

import "fmt"
import "runtime"
var print = fmt.Println

func x() {
  pc, file, line, _ := runtime.Caller(1)
  callerName := runtime.FuncForPC(pc).Name()
  print("file=", file, "callerName=", callerName, "line=", line)

}

func main() {
  x()
}
