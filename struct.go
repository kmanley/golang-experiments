package main

import "fmt"

type Order struct {
	side    int8  // Buy or Sell
	size    int8  // quantity
	price   int32 // price in ticks
	trader  string
	orderID int32
}

func main() {
	//o := Order{size: 123, trader: "Kevin"}
	o := Order{123, 0, 0, "Tom", 0}
	fmt.Println(o)
}
