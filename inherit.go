package main

import "fmt"

func func1(x, y int) {
	fmt.Printf("func1, %d, %d\n", x, y)
}

func func2(x, y int) {
	fmt.Printf("func2, %d, %d\n", x, y)
}

func func3(x, y int) {
	fmt.Printf("func3, %d, %d\n", x, y)
}

func func4(x, y int) {
	fmt.Printf("func4, %d, %d\n", x, y)
}

type TemplateProvider interface {
	HomeTemplate() func(int, int)
	AboutTemplate() func(int, int)
}

type Base struct {
	TemplateProvider
}

func (b Base) HomeTemplate() func(int, int) {
	return func1
}

func (b Base) AboutTemplate() func(int, int) {
	return func2
}

type Sub struct {
	Base
}

func (s Sub) HomeTemplate() func(int, int) {
	return func3
}

func (s Sub) AboutTemplate() func(int, int) {
	return func3
}

func doit(t TemplateProvider, x, y int) {
	template := t.HomeTemplate()
	if template != nil {
		template(x, y)
	} else {
		fmt.Printf("null template!")
	}
}

func main() {
	//s := Base{}
	s := Sub{}
	doit(s, 5, 5)

}
