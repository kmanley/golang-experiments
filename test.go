package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("/home/kevin/Downloads/FFFFFF-0.0.png")
	fmt.Printf("%#v", bytes)
}
