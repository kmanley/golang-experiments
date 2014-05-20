package main

import (
	_ "container/list"
	"fmt"
	_ "github.com/kmanley/gorderbook"
)

type Hand struct {
}

func main() {
	slice1 := []int{0, 0, 0, 0, 0}
	slice2 := []int{0, 0, 0, 0, 0}
	fmt.Println(slice1 < slice2)

}

/*
func main() {
	l := list.
}

func _main() {
	// See "Trading and Exchanges" by Harris, p126 (Continuous Trading Example)
	book := gorderbook.NewOrderBook("BTCUSD", 0, 10000, gorderbook.LogExecute)
	(&book).LimitOrder(gorderbook.Buy, 3, 200, "Bea")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Sell, 2, 201, "Sam")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Buy, 2, 200, "Ben")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Sell, 1, 198, "Sol")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Sell, 5, 202, "Stu")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Buy, 4, 202, "Bif")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Buy, 2, 201, "Bob")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Sell, 6, 200, "Sue")
	(&book).Dump()
	(&book).LimitOrder(gorderbook.Buy, 7, 198, "Bud")
	(&book).Dump()
}
*/
