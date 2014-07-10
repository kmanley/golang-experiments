package main

import "fmt"
import "regexp"

func main() {
	s := ""
	regex, err := regexp.Compile(s)
	// compiling empty regex returns a non-nil compiled regex
	fmt.Println(regex == nil, err == nil)
	// which seems to match anything...
	fmt.Println(regex.MatchString("foo"))
}
