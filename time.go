package main

import (
	"fmt"
	"strings"
	"time"
)

var print = fmt.Println

func testAfterFunc() {
	x := 3
	y := 4
	adder := func() {
		print("x*y=", x*y)
	}
	d, _ := time.ParseDuration("5s")
	time.AfterFunc(d, adder)
}

func main() {
	now := time.Now()
	print(now.UTC().Unix() - 1402403952)
	print(now.UTC().UnixNano() - 1402403952000000000)
	ident := strings.Replace(now.UTC().Format("060102150405.999"), ".", "", 1)
	print(ident)

	print("----")
	for i := 0; i < 20; i++ {
		print(time.Now().UnixNano())
	}
	print("----")

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
