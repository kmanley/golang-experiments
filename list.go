package main

import "container/list"
import "fmt"

func test(name string) {
	fmt.Println(name)
}

func main() {
	fmt.Println("hi")
	l := list.New()
	l.PushBack(10)
	l.PushBack(100)
	l.PushBack(1000)
	l.PushBack(2000)
	fmt.Println(l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
	//test()
	fmt.Println("----")
	fmt.Println(l.Front().Value)

	l.Remove(l.Front())
	fmt.Println("----")
	fmt.Println(l.Front().Value)

}
