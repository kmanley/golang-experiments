package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("eraseme.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
