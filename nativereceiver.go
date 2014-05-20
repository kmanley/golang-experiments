package main

import "fmt"

var print = fmt.Println

// can't define methods with int as receiver; you have to use your own type
// cannot define new methods on non-local type int32
// func (n int32) next() int {
// 	return n + 1
// }

type MyInt int32

func (n MyInt) next() MyInt {
	return n + 1
}

func main() {
	// var x MyInt = 1 // works
	// x := MyInt{5} // doesn't work: invalid type for composite literal: MyInt
	// x := 5.(MyInt) // doesn't work
	x := MyInt(5) // works and is simplest
	// x := new(MyInt) // works, initializes to 0
	//x := new(MyInt) // this plus below also works
	//*x = 10 // needs line above

	print(x.next())
	//print(x.(int)) // TODO: fails--how do we get this back to a regular int?
	//var y int32 = x // TODO: this fails--why?
	//print(y)

}
