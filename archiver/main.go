package main

import "fmt"
import "github.com/mholt/archiver"

func main() {
	err := archiver.Zip.Make("output.zip", []string{"."})
	fmt.Println(err)
}
