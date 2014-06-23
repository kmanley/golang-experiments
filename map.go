package main

import "fmt"

var print = fmt.Println

func test1() {
	people := make(map[string]int)
	people["Kevin"] = 46
	people["Magnus"] = 6
	people["Thor"] = 4
	people["Amy"] = 42
	print(people)
	//print(people.keys())
	age, ok := people["Kevin"]
	print(age, ok)
	age, ok = people["Fred"]
	print(age, ok)
	if age, ok = people["Magnus"]; ok {
		print(age)
	}
	for key, val := range people {
		print(key, val)
	}
}

func testDeleteWhileIterating() {
	people := make(map[string]int)
	people["Amy"] = 42
	people["Kevin"] = 46
	people["Thor"] = 4
	people["Magnus"] = 6
	for key, value := range people {
		print(key, value)
		delete(people, key)
	}
	print("---")
	fmt.Println("map len is now", len(people))
	for key, value := range people {
		print(key, value)
	}
}

func main() {
	testDeleteWhileIterating()
}
