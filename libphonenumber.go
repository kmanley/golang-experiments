package main

import (
	"log"

	"github.com/ttacon/libphonenumber"
)

func main() {
	num, err := libphonenumber.Parse("20501-5006", "US")
	log.Print(num.GetNationalNumber(), err)
}
